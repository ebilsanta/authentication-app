package utils

import (
    "log"
    "os"

    "github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
	// "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/service/sts"
)

func GetAWSCredentials() (string, string, string, error) {
	var (
		access_key_id string
		secret_access_key string
		session_token string
	)
	err := godotenv.Load("otp.env")
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

		log.Println("Using role credentials")
		access_key_id = *stsResp.Credentials.AccessKeyId
		secret_access_key = *stsResp.Credentials.SecretAccessKey
		session_token = *stsResp.Credentials.SessionToken
    } else {
		log.Println("Using dev credentials")
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

func ConnectSES() (s *ses.SES) {
	access_key_id, secret_access_key, session_token, err := GetAWSCredentials()

	if err != nil {
		log.Println("Could not get credentials for SES")
		return nil
	}

	return ses.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: credentials.NewStaticCredentials(access_key_id, secret_access_key, session_token),
	})))
}