package main

import (
	"context"
	// "fmt"
	"log"
	"net"
	"net/http"
	// "os"

	"github.com/golang/glog"
  	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"

	authentication "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/api/proto"
)

type myAuthenticationServer struct {
	authentication.UnimplementedAuthenticationServer
}

func (s myAuthenticationServer) CheckUserEnrolled(c context.Context, user *authentication.UserEnrolledRequest) (*authentication.UserEnrolledResponse, error) {
	// usersCollection := configs.GetCollection(configs.DB, "users")
	// var result bson.M
	// err := usersCollection.FindOne(
	// 	context.TODO(),
	// 	bson.D{
	// 		{"email", user.Email},
	// 		{"first_name", user.FirstName},
	// 		{"last_name", user.LastName},
	// 		{"birthdate", user.Birthdate},
	// 	},
	// ).Decode(&result)
	// if err != nil {
	// 	// ErrNoDocuments means that the filter did not match any documents in
	// 	// the collection.
	// 	if err == mongo.ErrNoDocuments {
	// 		return &authentication.UserEnrolledResponse{
	// 			Message: "Sorry you are not enrolled in our database, please make sure that you have an account with one of our partners!",
	// 		}, nil
	// 	}
	// 	log.Fatal(err)
	// }
	// if result["status"] == "pending" {
	// 	return &authentication.UserEnrolledResponse{
	// 		Message: "You are enrolled in in our database, please proceed to register an account with us!",
	// 	}, nil
	// }
	return &authentication.UserEnrolledResponse{
		Message: "You are enrolled in in our database, please proceed to register an account with us!",
	}, nil
}

func (s myAuthenticationServer) Register(c context.Context, userCredentials *authentication.RegistrationRequest) (*authentication.RegistrationResponse, error) {
	// usersCollection := configs.GetCollection(configs.DB, "users")
	// var result bson.M
	// err := usersCollection.FindOne(
	// 	context.TODO(),
	// 	bson.D{
	// 		{"email", userCredentials.Email},
	// 	},
	// ).Decode(&result)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		return &authentication.RegistrationResponse{
	// 			Email: userCredentials.Email, 
	// 			Message: "Sorry you are not enrolled in our database, please make sure that you have an account with one of our partners!",
	// 		}, nil
	// 	}
	// 	log.Fatal(err)
	// }
	// if result["status"] != "pending" {
	// 	return &authentication.RegistrationResponse{
	// 		Email: userCredentials.Email, 
	// 		Message: "You have already registered, please proceed to login instead!",
	// 	}, nil
	// }
	
	return &authentication.RegistrationResponse{
		Email: userCredentials.Email, 
		Message: "Registration successful, email verification sent!",
	}, nil
}


func main() {
  	defer glog.Flush()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	authenticationService := &myAuthenticationServer{}
	authentication.RegisterAuthenticationServer(serverRegistrar, authenticationService)

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