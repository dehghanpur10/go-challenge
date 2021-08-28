package deviceService

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"go-challenge/models"
	"go-challenge/repository/dynamoDB"
	"log"
	"os"
)



//CreateCore is struct for handle request, dynamoDB client and marshalMap are dependency injection
type CreateCore struct {
	db dynamoDB.DeviceDynamoDB
}

//NewCreateService is function for create new core for handler lambada
func NewCreateService(db dynamoDB.DeviceDynamoDB) *CreateCore {
	return &CreateCore{
		db: db,
	}
}

//CreateDevice is a lambda for handle post request from api Getway
func (d *CreateCore) CreateDevice(entity models.Device) error {
	device, _ := dynamodbattribute.MarshalMap(entity)
	input := &dynamodb.PutItemInput{
		Item:      device,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}
	_, err := d.db.PutItem(input)
	if err != nil {
		log.Println(err)
		return errors.New("server error")
	}
	return nil
}
