package deviceService

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-challenge/mocks"
	"go-challenge/models"
	"testing"
)

func TestCreateService(t *testing.T) {

	tests := []struct {
		name          string
		putItemError  error
		errorExpected error
	}{
		{name: "ok"},
		{name:"service error",putItemError: errors.New("error"),errorExpected: errors.New("server error")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB := new(mocks.DeviceDynamoDB)
			mockDB.On("PutItem", mock.Anything).Return(&dynamodb.PutItemOutput{}, test.putItemError)
			service := NewCreateService(mockDB)

			err := service.CreateDevice(models.Device{})
			if err == nil {
				assert.Nil(t, test.errorExpected)
			} else {
				assert.Error(t, test.errorExpected, err.Error())
			}
		})
	}

}
