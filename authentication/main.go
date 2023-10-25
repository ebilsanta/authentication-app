package main

import (
	"context"
	// "fmt"
	"log"
	"net"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	utils "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/utils"
	authenticationRepo "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/repository"
	authenticationUsecase "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/usecase"
	handler "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/handler/grpc"
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

	repo := authenticationRepo.NewDynamoDBAuthenticationRepository(db)
	usecase := authenticationUsecase.NewAuthenticationUsecase(repo)


	serverRegistrar := grpc.NewServer()
	handler.NewAuthenticationServer(serverRegistrar, usecase)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}