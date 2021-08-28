package dynamoDB

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type DeviceDynamoDB interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}

func NewDeviceDB(db dynamodbiface.DynamoDBAPI) DeviceDynamoDB {
	return &deviceDynamoDB{
		db: db,
	}
}

type deviceDynamoDB struct {
	db dynamodbiface.DynamoDBAPI
}

func (d *deviceDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	item, err := d.db.PutItem(input)
	return item, err
}

func (d *deviceDynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	item, err := d.db.GetItem(input)
	return item, err
}
