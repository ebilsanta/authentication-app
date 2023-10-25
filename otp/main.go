package main

import (
	"context"
	"log"
	"net"
	// "os"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	utils "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal/utils"
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

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}