package main

import (
	"context"
	"log"
	"net"
	"net/http"
	// "net/smtp"
	// "os"

	"github.com/golang/glog"
  	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"

	utils "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/utils"
	otp "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/api/proto"
	otpUsecase "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/usecase"
	otpRepo "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/repository"
	handler "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/handler/grpc"
)


func main() {
	defer glog.Flush()

	db := utils.GetDB()
	if db == nil {
		log.Fatalf("cannot connect to DB")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	repo := otpRepo.NewDynamoDBOTPRepository(db)
	usecase := otpUsecase.NewOTPUsecase(repo)
	handler.NewOTPServer(serverRegistrar, usecase)

	log.Println("Serving gRPC on 0.0.0.0:8089")
	go func() {
		log.Fatalln(serverRegistrar.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8089",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = otp.RegisterOTPHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8080")
	log.Fatalln(gwServer.ListenAndServe())
}