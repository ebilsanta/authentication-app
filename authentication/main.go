package main

import (
	"bytes"
	"context"
	// "fmt"
	"os"
	"log"
	"net"
	"encoding/json"
	// "io/ioutil"
	"net/http"
	"net/url"


	"github.com/golang/glog"
	"google.golang.org/grpc"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"

	authentication "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/api/proto"
	utils "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/utils"
	authenticationRepo "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/repository"
	authenticationUsecase "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/usecase"
	handler "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal/handler/grpc"
)

func processMessageBody(body string) map[string]string {
	data := map[string]string{}
	json.Unmarshal([]byte(body), &data)
	return data
}

func pollSQS() {
	queue := utils.ConnectSQS()
    for {
        // Poll SQS queue for messages
        messages, err := queue.ReceiveMessage(&sqs.ReceiveMessageInput{
            QueueUrl:            aws.String(os.Getenv("SQS_REQUEST_QUEUE_URL")),
            MaxNumberOfMessages: aws.Int64(10),
            WaitTimeSeconds:     aws.Int64(0),
			MessageAttributeNames: []*string{aws.String("All")},
        })
        if err != nil {
            log.Println("Error receiving message:", err)
            continue
        }

        // Process received messages
        for _, message := range messages.Messages {
			log.Print(message)
            // Forward message as gRPC request
            go handleMessage(message)
        }
    }
}

func handleMessage(message *sqs.Message) {
    // Extract data from message
	path := *message.MessageAttributes["Path"].StringValue
	data := processMessageBody(*message.Body)
	callback := data["callback"]

    // Create gRPC client
    conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
    if err != nil {
        log.Println("Error connecting to gRPC server:", err)
        return
    }
    defer conn.Close()

    // Create gRPC client
    client := authentication.NewAuthenticationClient(conn)

	if path == "register" {
		response, err := client.Register(context.Background(), &authentication.RegisterRequest{
			Company: data["company"],
			Email: data["email"],
			FirstName: data["firstName"],
			LastName: data["lastName"],
			Birthdate: data["birthdate"],
			Password: data["password"],
		})
		if err != nil {
			log.Println("Error making gRPC request:", err)
			return
		}
		input := map[string]string{
			"email": response.Email,
			"message": response.Message,
			"verification_key": response.VerificationKey,
		}
		err = ResponseMessage(input, callback, *message)
		if err != nil {
			log.Println("Error returning response:", err)
			return
		}
		return
	} else if path == "verify-email" {
		response, err := client.VerifyEmail(context.Background(), &authentication.VerifyEmailRequest{
			VerificationKey: data["verificationKey"],
			Otp: data["otp"],
			Email: data["email"],
		})
		if err != nil {
			log.Println("Error making gRPC request:", err)
			return
		}
		input := map[string]string{
			"status": response.Status,
			"details": response.Details,
			"email": response.Email,
		}
		err = ResponseMessage(input, callback, *message)
		if err != nil {
			log.Println("Error returning response:", err)
			return
		}
		return
	} else if path == "login" {
		response, err := client.Login(context.Background(), &authentication.LoginRequest{
			Company: data["company"],
			Email: data["email"],
			Password: data["password"],
		})
		if err != nil {
			log.Println("Error making gRPC request:", err)
			return
		}
		input := map[string]string{
			"status": response.Status,
			"idToken": response.IdToken,
		}
		err = ResponseMessage(input, callback, *message)
		if err != nil {
			log.Println("Error returning response:", err)
			return
		}
		return
	} else if path == "otp" {
		response, err := client.NewOTP(context.Background(), &authentication.NewOTPRequest{
			Company: data["company"],
			Email: data["email"],
		})
		if err != nil {
			log.Println("Error making gRPC request:", err)
			return
		}
		input := map[string]string{
			"message": response.Message,
			"verification_key": response.VerificationKey,
		}
		err = ResponseMessage(input, callback, *message)
		if err != nil {
			log.Println("Error returning response:", err)
			return
		}
		return
	} else if path == "valid-token" {
		response, err := client.ValidToken(context.Background(), &authentication.ValidTokenRequest{
			VerificationKey: data["verificationKey"],
			Otp: data["otp"],
			Email: data["email"],
		})
		if err != nil {
			log.Println("Error making gRPC request:", err)
			return
		}
		input := map[string]string{
			"message": response.Message,
			"token": response.Token,
		}
		err = ResponseMessage(input, callback, *message)
		if err != nil {
			log.Println("Error returning response:", err)
			return
		}
		return
	} else if path == "change-password" {
		response, err := client.ChangePassword(context.Background(), &authentication.ChangePasswordRequest{
			Token: data["token"],
			Company: data["company"],
			Email: data["email"],
			Password: data["password"],
		})
		if err != nil {
			log.Println("Error making gRPC request:", err)
			return
		}
		input := map[string]string{
			"status": response.Status,
			"details": response.Details,
			"email": response.Email,
		}
		err = ResponseMessage(input, callback, *message)
		if err != nil {
			log.Println("Error returning response:", err)
			return
		}
		return
	} else {
		log.Println("Invalid route!")
		return
	}
}

func ResponseMessage(input map[string]string, callback string, message sqs.Message) (error) {
	queue := utils.ConnectSQS()
	postBody, _ := json.Marshal(input)
	responseBody := bytes.NewBuffer(postBody)
	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(os.Getenv("SQS_REQUEST_QUEUE_URL")),
		ReceiptHandle: message.ReceiptHandle,
	}
	_, err := url.ParseRequestURI(callback)
	if err != nil {
		log.Println("Error with callback URL: ", err)
		_, err = queue.DeleteMessage(deleteParams)
		if err != nil {
			log.Println(err)
		}
		return nil
	}
	resp, err := http.Post(callback, "application/json", responseBody)
	if err != nil {
		log.Println("An Error Occured: ", err)
		_, err = queue.DeleteMessage(deleteParams)
		if err != nil {
			log.Println(err)
		}
		return nil
	}
	_, err = queue.DeleteMessage(deleteParams)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	return nil
}


func main() {
  	defer glog.Flush()

	db := utils.ConnectDB()
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

	for i := 1; i <= 10; i++ {
		go pollSQS()
	}

	func() {
		log.Fatalln(serverRegistrar.Serve(lis))
	}()
}