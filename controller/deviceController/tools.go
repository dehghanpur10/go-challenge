package deviceController

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go-challenge/models"
	"log"
	"net/http"
	"os"
)

func GetDynamoDB() (*dynamodb.DynamoDB, error) {
	region := os.Getenv("AWS_REGION")
	accessToken := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
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

func CreateError(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	result, _ := json.Marshal(models.Error{
		Message: err,
	})
	_, _ = w.Write(result)
}