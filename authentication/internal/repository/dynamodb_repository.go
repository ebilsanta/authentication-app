package repository

import (
	"os"
	"log"
	// "fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	models "github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/authentication/internal"
)

type DynamoDBAuthenticationRepository struct {
	DB *dynamodb.DynamoDB
}

func NewDynamoDBAuthenticationRepository(DB *dynamodb.DynamoDB) AuthenticationRepository {
	return &DynamoDBAuthenticationRepository{DB}
}

func (d *DynamoDBAuthenticationRepository) GetUserByFullInfo(company string, email string, first_name string, last_name string, birthdate string) (*models.User, error) {
	result, err := d.DB.Query(&dynamodb.QueryInput{
		TableName: aws.String(os.Getenv("USER_TABLE")),
		KeyConditionExpression: aws.String("company = :company AND email = :email"),
		FilterExpression: aws.String("first_name = :first_name AND last_name = :last_name AND birthdate = :birthdate"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":company": {
				S: aws.String(company),
			},
			":email": {
				S: aws.String(email),
			},
			":first_name": {
				S: aws.String(first_name),
			},
			":last_name": {
				S: aws.String(last_name),
			},
			":birthdate": {
				S: aws.String(birthdate),
			},
		},
	})

	if err != nil {
		log.Printf("Couldn't execute query to get user. Here's why: %v\n", err)
		return nil, err
	} else if len(result.Items) == 0 {
		return nil, nil
	}

	var user models.User
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &user)
	if err != nil {
		log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
	}

	return &user, nil
}

func (d *DynamoDBAuthenticationRepository) UpdateUserByEmail(company string, email string) (*models.User, error) {
	status := "status"
	
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("USER_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"company": {
				S: aws.String(company),
			},
			"email": {
				S: aws.String(email),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":status": {
				S: aws.String("Verified"),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#status": &status,
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #status = :status"),
	}
	
	output, err := d.DB.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

	var updated models.User

	err = dynamodbattribute.UnmarshalMap(output.Attributes, &updated)
	if err != nil {
		log.Printf("Couldn't unmarshall update response. Here's why: %v\n", err)
	}

	return &updated, nil
}

func (d *DynamoDBAuthenticationRepository) GetCredentialByEmail(company string, email string) (*models.Credential, error) {
	result, err := d.DB.Query(&dynamodb.QueryInput{
		TableName: aws.String(os.Getenv("CREDENTIAL_TABLE")),
		KeyConditionExpression: aws.String("company = :company AND email = :email"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":company": {
				S: aws.String(company),
			},
			":email": {
				S: aws.String(email),
			},
		},
	})

	if err != nil {
		log.Printf("Couldn't execute query to get user. Here's why: %v\n", err)
		return nil, err
	} else if len(result.Items) == 0 {
		return nil, nil
	}

	var credential models.Credential
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &credential)
	if err != nil {
		log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
	}

	return &credential, nil
}

func (d *DynamoDBAuthenticationRepository) GetUserByEmail(company string, email string) (*models.User, error) {
	result, err := d.DB.Query(&dynamodb.QueryInput{
		TableName: aws.String(os.Getenv("USER_TABLE")),
		KeyConditionExpression: aws.String("company = :company AND email = :email"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":company": {
				S: aws.String(company),
			},
			":email": {
				S: aws.String(email),
			},
		},
	})

	if err != nil {
		log.Printf("Couldn't execute query to get user. Here's why: %v\n", err)
		return nil, err
	} else if len(result.Items) == 0 {
		return nil, nil
	}

	var user models.User
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &user)
	if err != nil {
		log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
	}

	return &user, nil
}

func (d *DynamoDBAuthenticationRepository) DeleteCredentialByEmail(company string, email string) (string, error) {
	return "User Deleted", nil
}

func (d *DynamoDBAuthenticationRepository) RegisterUser(company string, email string, password string) (*models.Credential, error) {
	credential := &models.Credential{company, email, password}

	av, err := dynamodbattribute.MarshalMap(credential)
	if err != nil {
		log.Printf("Got error marshalling new movie item: %s", err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("CREDENTIAL_TABLE")),
	}
	
	_, err = d.DB.PutItem(input)
	if err != nil {
		log.Printf("Got error calling PutItem: %s", err)
		return nil, err
	}

	return credential, nil
}