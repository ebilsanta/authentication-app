package main

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	utils "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/utils"
	otp "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/api/proto"
	otpUsecase "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/usecase"
	otpRepo "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/repository"
	handler "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/handler/grpc"
)


func server(ctx context.Context) (otp.OTPClient, func()) {
	buffer := 1024 * 1024
	listener := bufconn.Listen(buffer)

	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	db := utils.ConnectDB()
	sesService := utils.ConnectSES()
	repo := otpRepo.NewDynamoDBOTPRepository(db)
	usecase := otpUsecase.NewOTPUsecase(repo, sesService)
	handler.NewOTPServer(baseServer, usecase)

	go func() {
		if err := baseServer.Serve(listener); err != nil {
			log.Println("err", err)
		}
	}()

	conn, _ := grpc.DialContext(ctx, "", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())

	closer := func() {
		listener.Close()
		baseServer.Stop()
	}

	client := otp.NewOTPClient(conn)

	return client, closer
}

func TestOTPService(t *testing.T) {
	client, closer := server(context.Background())
	defer closer()

	log.Println("Starting GetOTP OK Test")
	otp_res, err := client.GetOTP(context.Background(), &otp.OTPRequest{Company: "ascenda", Email: "cs301.auth.otp@gmail.com"})
	if err != nil {
		t.Fatalf("client.GetOTP %v", err)
	}

	if otp_res.VerificationKey == "" || otp_res.Message != "Success" {
		t.Fatalf("Unexpected values %v", otp_res)
	}

	verification_key := otp_res.VerificationKey

	log.Println("Starting GetOTP Invalid Email Test")
	_, err = client.GetOTP(context.Background(), &otp.OTPRequest{Company: "ascenda", Email: "cs301.auth.otp@gmail.co"})
	if err == nil {
		t.Fatalf("Invalid Email was not detected")
	}

	log.Println("Starting VerifyOTP OTP Not Matching Test")
	verify_res, err := client.VerifyOTP(context.Background(), &otp.VerifyOTPRequest{VerificationKey: verification_key, Otp: "", Email: "cs301.auth.otp@gmail.com"})
	if err != nil {
		t.Fatalf("client.VerifyOTP %v", err)
	}

	if verify_res.Status != "Failure" || verify_res.Details != "Some details did not match"{
		t.Fatalf("Unexpected values %v", verify_res)
	}

	log.Println("Starting VerifyOTP Invalid Verification Key Test")
	verify_res, err = client.VerifyOTP(context.Background(), &otp.VerifyOTPRequest{VerificationKey: "", Otp: "", Email: "cs301.auth.otp@gmail.com"})
	if err != nil {
		t.Fatalf("client.VerifyOTP %v", err)
	}

	if verify_res.Status != "Failure" || verify_res.Details != "Error getting details from verification key"{
		t.Fatalf("Unexpected values %v", verify_res)
	}
}