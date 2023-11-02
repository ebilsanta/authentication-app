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

func GetAWSCredentials() (string, string, string, error) {
	var (
		access_key_id string
		secret_access_key string
		session_token string
	)
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
			return "", "", "", err
		}

		access_key_id = *stsResp.Credentials.AccessKeyId
		secret_access_key = *stsResp.Credentials.SecretAccessKey
		session_token = *stsResp.Credentials.SessionToken
    } else {
		access_key_id = os.Getenv("AWS_ACCESS_KEY_ID")
		secret_access_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
		session_token = ""
	}
	return access_key_id, secret_access_key, session_token, nil
}

func ConnectDB() (db *dynamodb.DynamoDB) {
	access_key_id, secret_access_key, session_token, err := GetAWSCredentials()

	if err != nil {
		log.Println("Could not get credentials for DynamoDB")
		return nil
	}

	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: credentials.NewStaticCredentials(access_key_id, secret_access_key, session_token),
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
    access_key_id, secret_access_key, session_token, err := GetAWSCredentials()

	if err != nil {
		log.Println("Could not get credentials for SQS")
		return nil
	}

	return sqs.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: credentials.NewStaticCredentials(access_key_id, secret_access_key, session_token),
	})))
}

func GetPrivateKey() string {
	access_key_id, secret_access_key, session_token, err := GetAWSCredentials()

	if err != nil {
		log.Println("Could not get credentials for Secrets Manager")
		return ""
	}

	sm := secretsmanager.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: credentials.NewStaticCredentials(access_key_id, secret_access_key, session_token),
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