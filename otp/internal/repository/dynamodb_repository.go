package repository

import (
	"os"
	"log"
	"time"
	"strconv"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	models "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/otp/internal"
)

type DynamoDBOTPRepository struct {
	DB *dynamodb.DynamoDB
}

func NewDynamoDBOTPRepository(DB *dynamodb.DynamoDB) OTPRepository {
	return &DynamoDBOTPRepository{DB}
}

func (d *DynamoDBOTPRepository) CreateOTP(otp string, expiration_date time.Time) (*models.OTP, error) {
	otp_num, _ := strconv.Atoi(otp)
	OTP := &models.OTP{otp_num, expiration_date, false}

	av, err := dynamodbattribute.MarshalMap(OTP)
	if err != nil {
		log.Printf("Got error marshalling new movie item: %s", err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("OTP_TABLE")),
	}

	fmt.Println(input)
	
	_, err = d.DB.PutItem(input)
	if err != nil {
		log.Printf("Got error calling PutItem: %s", err)
		return nil, err
	}

	return OTP, nil
}

func (d *DynamoDBOTPRepository) GetOTP(otp string, expiration_date time.Time) (*models.OTP, error) {
	result, err := d.DB.Query(&dynamodb.QueryInput{
		TableName: aws.String(os.Getenv("OTP_TABLE")),
		KeyConditionExpression: aws.String("expiration_date = :expiration_date AND otp = :otp"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":expiration_date": {
				S: aws.String(expiration_date.String()),
			},
			":otp": {
				N: aws.String(string(otp)),
			},
		},
	})

	if err != nil {
		log.Printf("Couldn't execute query to get user. Here's why: %v\n", err)
		return nil, err
	} else if len(result.Items) == 0 {
		return nil, nil
	}

	var OTP models.OTP
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &OTP)
	if err != nil {
		log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
	}

	return &OTP, nil
}

func (d *DynamoDBOTPRepository) UpdateOTP(otp string, expiration_date time.Time) (*models.OTP, error) {

	verified := true

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("OTP_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"otp": {
				N: aws.String(string(otp)),
			},
			"expiration_date": {
				S: aws.String(expiration_date.String()),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":verified": {
				BOOL: &verified,
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set verified = :verified"),
	}
	
	output, err := d.DB.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

	var updated models.OTP

	err = dynamodbattribute.UnmarshalMap(output.Attributes, &updated)
	if err != nil {
		log.Printf("Couldn't unmarshall update response. Here's why: %v\n", err)
	}

	return &updated, nil
}