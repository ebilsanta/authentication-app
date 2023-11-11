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
	// "github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
)

func GetAWSCredentials() credentials.Credentials {
	err := godotenv.Load("authentication.env")
    if err != nil {
		role := os.Getenv("ROLE_ARN")
		sess := session.Must(session.NewSession())
		creds := stscreds.NewCredentials(sess, role)

		log.Println("Using role credentials")
		return *creds
    }
	log.Println("Using dev credentials")
	return *credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
}

func ConnectDB() (db *dynamodb.DynamoDB) {
	sess := session.Must(session.NewSession())
	creds := GetAWSCredentials()

	return dynamodb.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: &creds,
	})
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
	sess := session.Must(session.NewSession())
    creds := GetAWSCredentials()

	return sqs.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: &creds,
	})
}

func GetPrivateKey() string {
	sess := session.Must(session.NewSession())
	creds := GetAWSCredentials()

	sm := secretsmanager.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: &creds,
	})
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

func GetPublicKey() string {
	sess := session.Must(session.NewSession())
	creds := GetAWSCredentials()

	sm := secretsmanager.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: &creds,
	})
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

	return data["PUB_KEY"]
}