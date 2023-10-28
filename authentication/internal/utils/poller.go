package utils

// import (
// 	"fmt"
// 	"bytes"
// 	"context"
// 	"log"
// 	"io/ioutil"
//    	"net/http"
// 	"encoding/json"
// 	"os"

// 	"google.golang.org/grpc"
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/service/sqs"

// 	authentication "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/api/proto"
// )

// func pollSQS(chn chan<- *sqs.Message) {
// 	queue := GetSQS()
//     for {
//         // Poll SQS queue for messages
//         messages, err := queue.ReceiveMessage(&sqs.ReceiveMessageInput{
//             QueueUrl:            aws.String(os.Getenv("SQS_REQUEST_QUEUE_URL")),
//             MaxNumberOfMessages: aws.Int64(10),
//             WaitTimeSeconds:     aws.Int64(20),
//         })
//         if err != nil {
//             log.Println("Error receiving message:", err)
//             continue
//         }

//         // Process received messages
//         for _, message := range messages.Messages {
//             // Forward message as gRPC request
// 			fmt.Println(message)
//             chn <- message
//         }
//     }
// }

// func handleMessage(message *sqs.Message) {
//     // Extract data from message
//     path := *message.Body
// 	callback := *message.Attributes["Callback"]

//     // Create gRPC client
//     conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
//     if err != nil {
//         log.Println("Error connecting to gRPC server:", err)
//         return
//     }
//     defer conn.Close()

//     // Create gRPC client
//     client := authentication.NewAuthenticationClient(conn)

// 	if path == "/register" {
// 		response, err := client.Register(context.Background(), &authentication.RegisterRequest{
// 			Company: *message.Attributes["Company"],
// 			Email: *message.Attributes["Email"],
// 			FirstName: *message.Attributes["FirstName"],
// 			LastName: *message.Attributes["LastName"],
// 			Birthdate: *message.Attributes["Birthdate"],
// 			Password: *message.Attributes["Password"],
// 		})
// 		if err != nil {
// 			log.Println("Error making gRPC request:", err)
// 			return
// 		}
// 		input := map[string]string{
// 			"email": response.Email,
// 			"message": response.Message,
// 			"verification_key": response.VerificationKey,
// 		}
// 		err = ResponseMessage(input, callback, *message)
// 		if err != nil {
// 			log.Println("Error returning response:", err)
// 			return
// 		}
// 		return
// 	} else if path == "/verify" {
// 		response, err := client.VerifyEmail(context.Background(), &authentication.VerifyEmailRequest{
// 			VerificationKey: *message.Attributes["VerificationKey"],
// 			Otp: *message.Attributes["OTP"],
// 			Email: *message.Attributes["Email"],
// 		})
// 		if err != nil {
// 			log.Println("Error making gRPC request:", err)
// 			return
// 		}
// 		input := map[string]string{
// 			"status": response.Status,
// 			"details": response.Details,
// 			"email": response.Email,
// 		}
// 		err = ResponseMessage(input, callback, *message)
// 		if err != nil {
// 			log.Println("Error returning response:", err)
// 			return
// 		}
// 		return
// 	} else {
// 		log.Println("Invalid route!")
// 		return
// 	}
// }

// func ResponseMessage(input map[string]string, callback string, message sqs.Message) (error) {
// 	queue := GetSQS()
// 	fmt.Println(input)
// 	postBody, _ := json.Marshal(input)
// 	responseBody := bytes.NewBuffer(postBody)
// 	resp, err := http.Post(callback, "application/json", responseBody)
// 	if err != nil {
// 		log.Fatalf("An Error Occured %v", err)
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 		return err
// 	}
// 	sb := string(body)
// 	log.Printf(sb)
// 	deleteParams := &sqs.DeleteMessageInput{
// 		QueueUrl:      aws.String(os.Getenv("SQS_REQUEST_QUEUE_URL")),
// 		ReceiptHandle: message.ReceiptHandle,
// 	}
// 	_, err = queue.DeleteMessage(deleteParams)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	return nil
// }