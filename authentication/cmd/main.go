package main

import (
	"context"
	// "fmt"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
  	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"

	utils "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/utils"
	authentication "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/api/proto"
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
	err = authentication.RegisterAuthenticationHandler(context.Background(), gwmux, conn)
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