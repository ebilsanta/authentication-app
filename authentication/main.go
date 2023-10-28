package main

import (
	"bytes"
	"context"
	// "fmt"
	"os"
	"log"
	"net"
	"encoding/json"
	"io/ioutil"
	"net/http"


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

func pollSQS() {
	queue := utils.ConnectSQS()
    for {
        // Poll SQS queue for messages
        messages, err := queue.ReceiveMessage(&sqs.ReceiveMessageInput{
            QueueUrl:            aws.String(os.Getenv("SQS_REQUEST_QUEUE_URL")),
            MaxNumberOfMessages: aws.Int64(10),
            WaitTimeSeconds:     aws.Int64(20),
			MessageAttributeNames: []*string{aws.String("All")},
        })
        if err != nil {
            log.Println("Error receiving message:", err)
            continue
        }

        // Process received messages
        for _, message := range messages.Messages {
            // Forward message as gRPC request
            handleMessage(message)
        }
    }
}

func handleMessage(message *sqs.Message) {
    // Extract data from message
    path := *message.Body
	callback := *message.MessageAttributes["Callback"].StringValue

    // Create gRPC client
    conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
    if err != nil {
        log.Println("Error connecting to gRPC server:", err)
        return
    }
    defer conn.Close()

    // Create gRPC client
    client := authentication.NewAuthenticationClient(conn)

	if path == "/register" {
		response, err := client.Register(context.Background(), &authentication.RegisterRequest{
			Company: *message.MessageAttributes["Company"].StringValue,
			Email: *message.MessageAttributes["Email"].StringValue,
			FirstName: *message.MessageAttributes["FirstName"].StringValue,
			LastName: *message.MessageAttributes["LastName"].StringValue,
			Birthdate: *message.MessageAttributes["Birthdate"].StringValue,
			Password: *message.MessageAttributes["Password"].StringValue,
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
	} else if path == "/verify" {
		response, err := client.VerifyEmail(context.Background(), &authentication.VerifyEmailRequest{
			VerificationKey: *message.Attributes["VerificationKey"],
			Otp: *message.Attributes["OTP"],
			Email: *message.Attributes["Email"],
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
	resp, err := http.Post(callback, "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	sb := string(body)
	log.Printf(sb)
	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(os.Getenv("SQS_REQUEST_QUEUE_URL")),
		ReceiptHandle: message.ReceiptHandle,
	}
	_, err = queue.DeleteMessage(deleteParams)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}


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

	go pollSQS()

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}