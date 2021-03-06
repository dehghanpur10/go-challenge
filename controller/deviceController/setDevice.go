package deviceController

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-challenge/models"
	"go-challenge/repository/dynamoDB"
	"go-challenge/services/deviceService"
	"log"
	"net/http"
)

// SetDevice
// @Summary create a new device info
// @Description this endpoint create a new device info
// @Tags device
// @Accept  json
// @Produce  json
// @Param device body models.Device true "device id"
// @Success 201 {object} models.Device
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /devices [post]
func SetDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var device models.Device
	_ = json.NewDecoder(r.Body).Decode(&device)

	validate := validator.New()
	err := validate.Struct(device)
	if err != nil {
		log.Println(err)
		CreateError(w, "invalid device info", http.StatusBadRequest)
		return
	}

	db, err := GetDynamoDB()
	if err != nil {
		log.Println(err)
		CreateError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	service := deviceService.NewCreateService(dynamoDB.NewDeviceDB(db))
	err = service.CreateDevice(device)
	if err != nil {
		log.Println(err)
		CreateError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	result, _ := json.Marshal(device)
	_, _ = w.Write(result)

}

