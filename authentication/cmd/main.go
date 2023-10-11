package main

import (
	"context"
	authentication "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/api/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	// "os"
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
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	authenticationService := &myAuthenticationServer{}
	authentication.RegisterAuthenticationServer(serverRegistrar, authenticationService)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}