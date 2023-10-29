package utils

import (
	"context"
    "log"
    "os"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"

    "github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ConnectDB() (db *dynamodb.DynamoDB) {
    err := godotenv.Load("infrastructure/.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})))
}

var ddb *dynamodb.DynamoDB = ConnectDB()

func GetDB() (db *dynamodb.DynamoDB) {
	return ddb
}

func ConnectOTPServer() (*grpc.ClientConn, error) {
	err := godotenv.Load("authentication.env")
    if err != nil {
        log.Fatal("Error loading .env file")
		return nil, err
    }

	conn, err := grpc.DialContext(
		context.Background(),
		os.Getenv("OTP_URL"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	return conn, err
}

func ConnectSQS() (queue *sqs.SQS) {
    err := godotenv.Load("infrastructure/.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	return sqs.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})))
}

var queue *sqs.SQS = ConnectSQS()

func GetSQS() (queue *sqs.SQS) {
	return queue
}