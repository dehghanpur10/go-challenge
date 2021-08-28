package deviceController

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"os"
)

func GetDynamoDB() (*dynamodb.DynamoDB, error) {
	region := os.Getenv("AWS_REGION")
	accessToken := os.Getenv("ACCESS_TOKEN")
	secretKey := os.Getenv("SECRET_KEY")
	credential := credentials.NewStaticCredentials(accessToken, secretKey, "")
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credential,
	},

	)
	if err != nil {
		log.Println(err)
		return &dynamodb.DynamoDB{}, err
	}
	return dynamodb.New(awsSession), nil
}
