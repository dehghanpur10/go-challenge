package deviceController

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/go-playground/validator/v10"
	"go-challenge/models"
	"go-challenge/services/deviceService"
	"log"
	"net/http"
)

//SetDevice is controller for set device to dynamoDB
func SetDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var device models.Device
	_ = json.NewDecoder(r.Body).Decode(&device)

	validate := validator.New()
	err := validate.Struct(device)
	if err != nil {
		log.Println(err)
		createError(w, "invalid device info", http.StatusBadRequest)
		return
	}

	db, err := GetDynamoDB()
	if err != nil {
		log.Println(err)
		createError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	service := deviceService.NewCreateService(db, dynamodbattribute.MarshalMap)
	err = service.CreateDevice(device)
	if err != nil {
		log.Println(err)
		createError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	result, _ := json.Marshal(device)
	_, _ = w.Write(result)

}
func createError(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	result, _ := json.Marshal(models.Error{
		Message: err,
	})
	_, _ = w.Write(result)
}
