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
	// "github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
)

func GetAWSCredentials() credentials.Credentials {
	err := godotenv.Load("otp.env")
    if err != nil {
		role := os.Getenv("ROLE_ARN")
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		}))
		creds := stscreds.NewCredentials(sess, role)

		log.Println("Using role credentials")
		return *creds
    }
	log.Println("Using dev credentials")
	return *credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
}

func ConnectDB() (db *dynamodb.DynamoDB) {
	creds := GetAWSCredentials()

	return dynamodb.New(session.Must(session.NewSession()), &aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: &creds,
	})
}

func ConnectSES() (s *ses.SES) {
	creds := GetAWSCredentials()

	return ses.New(session.Must(session.NewSession()), &aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: &creds,
	})
}