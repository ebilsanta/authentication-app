import boto3
from botocore.exceptions import NoCredentialsError, ClientError
from os import getenv
from dotenv import load_dotenv

try:
    load_dotenv()
    dynamodb = boto3.resource('dynamodb', 
                            aws_access_key_id=getenv("AWS_ACCESS_KEY_ID"),
                            aws_secret_access_key=getenv("AWS_SECRET_ACCESS_KEY"),
                            region_name="ap-southeast-1",)
    table = dynamodb.Table('users')
except NoCredentialsError:
    print("Credentials not available")
    exit()

def get_user(company: str, email: str):
    try:
        reponse = table.get_item(
            Key={
                "company": company,
                "email": email
            }
        )
    except ClientError as e:
        print(e.response['Error']['Message'])
    else:
        return reponse['Item']