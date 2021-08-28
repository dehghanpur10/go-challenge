package deviceService

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go-challenge/models"
	"go-challenge/repository/dynamoDB"
	"log"
	"os"
)

//MarshalType is type for marshal function
type MarshalType func(in interface{}) (map[string]*dynamodb.AttributeValue, error)

//Core is struct for handle request, dynamoDB client and marshalMap are dependency injection
type Core struct {
	db         dynamoDB.DeviceDynamoDB
	marshalMap MarshalType
}

//NewCreateService is function for create new core for handler lambada
func NewCreateService(db dynamoDB.DeviceDynamoDB, marshal MarshalType) *Core {

	return &Core{
		db:         db,
		marshalMap: marshal,
	}
}

//CreateDevice is a lambda for handle post request from api Getway
func (d *Core) CreateDevice(entity models.Device) error {
	device, err := d.marshalMap(entity)
	if err != nil {
		log.Println(err)
		return errors.New("server error")
	}
	input := &dynamodb.PutItemInput{
		Item:      device,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}
	_, err = d.db.PutItem(input)
	if err != nil {
		log.Println(err)
		return errors.New("server error")
	}
	return nil
}
