package utils

import (
	"context"
	"encoding/json"
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
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func ConnectDB() (db *dynamodb.DynamoDB) {
	aws_session := ""
    err := godotenv.Load("authentication.env")
    if err != nil {
		log.Print("Development env file not found, trying production env")
		role := os.Getenv("ROLE_ARN")
		stsClient := sts.New(session.Must(session.NewSession()))
		params := &sts.AssumeRoleInput{
			RoleArn:         aws.String(role),
			RoleSessionName: aws.String("RoleSession0"),
		}
		stsResp, err := stsClient.AssumeRole(params)

		if err != nil {
			log.Println("Error getting role credentials: ", err)
			return nil
		}

		os.Setenv("AWS_ACCESS_KEY_ID", *stsResp.Credentials.AccessKeyId)
		os.Setenv("AWS_SECRET_ACCESS_KEY", *stsResp.Credentials.SecretAccessKey)
		os.Setenv("AWS_SESSION_TOKEN", *stsResp.Credentials.SessionToken)

		aws_session = *stsResp.Credentials.SessionToken  
    }
	log.Print("Using dev credentials")

	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), aws_session),
	})))
}

func ConnectOTPServer() (*grpc.ClientConn, error) {
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
    aws_session := ""
	err := godotenv.Load("authentication.env")
    if err != nil {
		log.Print("Development env file not found, trying production env")
		role := os.Getenv("ROLE_ARN")
		stsClient := sts.New(session.Must(session.NewSession()))
		params := &sts.AssumeRoleInput{
			RoleArn:         aws.String(role),
			RoleSessionName: aws.String("RoleSession0"),
		}
		stsResp, err := stsClient.AssumeRole(params)

		if err != nil {
			log.Println("Error getting role credentials: ", err)
			return nil
		}

		os.Setenv("AWS_ACCESS_KEY_ID", *stsResp.Credentials.AccessKeyId)
		os.Setenv("AWS_SECRET_ACCESS_KEY", *stsResp.Credentials.SecretAccessKey)
		os.Setenv("AWS_SESSION_TOKEN", *stsResp.Credentials.SessionToken)

		aws_session = *stsResp.Credentials.SessionToken  
    }

	return sqs.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), aws_session),
	})))
}

func GetPrivateKey() string {
	aws_session := ""
	err := godotenv.Load("authentication.env")
    if err != nil {
		log.Print("Development env file not found, trying production env")
		role := os.Getenv("ROLE_ARN")
		stsClient := sts.New(session.Must(session.NewSession()))
		params := &sts.AssumeRoleInput{
			RoleArn:         aws.String(role),
			RoleSessionName: aws.String("RoleSession0"),
		}
		stsResp, err := stsClient.AssumeRole(params)

		if err != nil {
			log.Println("Error getting role credentials: ", err)
			return ""
		}

		os.Setenv("AWS_ACCESS_KEY_ID", *stsResp.Credentials.AccessKeyId)
		os.Setenv("AWS_SECRET_ACCESS_KEY", *stsResp.Credentials.SecretAccessKey)
		os.Setenv("AWS_SESSION_TOKEN", *stsResp.Credentials.SessionToken)

		aws_session = *stsResp.Credentials.SessionToken  
    }

	sm := secretsmanager.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), aws_session),
	})))
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(os.Getenv("KEYS_NAME")),
	}
	result, err := sm.GetSecretValue(input)
	if err != nil {
		log.Println("Error getting private key: ", err)
		return ""
	}

	data := map[string]string{}
	json.Unmarshal([]byte(*result.SecretString), &data)

	return data["PVT_KEY"]
}