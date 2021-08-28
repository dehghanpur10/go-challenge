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

//NewGetService is function for create new core for handler lambada
func NewGetService(db dynamoDB.DeviceDynamoDB) *GetCore {
	return &GetCore{
		db: db,
	}
}

//GetCore is struct for handle request, dynamoDB client and marshalMap are dependency injection
type GetCore struct {
	db dynamoDB.DeviceDynamoDB
}

//GetDevice is a lambda for handle post request from api Getway
func (d *GetCore) GetDevice(id string) (models.Device, error) {
	key := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(id),
		},
	}
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key:       key,
	}
	result, err := d.db.GetItem(getItemInput)
	if err != nil {
		log.Println(err)
		return models.Device{}, errors.New("server error")
	}
	if result.Item == nil {
		log.Println(err)
		return models.Device{}, errors.New("device not found")
	}
	var device models.Device
	_ = dynamodbattribute.UnmarshalMap(result.Item, &device)
	return device, nil
}
