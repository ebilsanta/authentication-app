package utils

import (
	"context"
    "log"
    "os"

    "github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go-v2/config"
)

func ConnectDB() (db *dynamodb.DynamoDB) {
    err := godotenv.Load("otp.env")
    if err != nil {
		log.Print("Development env file not found, trying production env")
		config.LoadDefaultConfig(context.TODO())
		// err = godotenv.Load()
		// if err != nil {
		// 	log.Fatal("Error loading .env file")
		// }     
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

func ConnectSES() (s *ses.SES) {
    // err := godotenv.Load("otp.env")
    // if err != nil {
	// 	log.Print("Development env file not found, trying production env")
	// 	err = godotenv.Load()
	// 	if err != nil {
	// 		log.Fatal("Error loading .env file")
	// 	}     
    // }

	return ses.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})))
}

var awsses *ses.SES = ConnectSES()

func GetSES() (s *ses.SES) {
	return awsses
}