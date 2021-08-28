package deviceController

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-challenge/repository/dynamoDB"
	"go-challenge/services/deviceService"
	"log"
	"net/http"
)

//GetDevice is controller for get device from dynamoDB
func GetDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	db, err := GetDynamoDB()
	if err != nil {
		log.Println(err)
		createError(w, "internal server error", http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)

	service := deviceService.NewGetCoreService(dynamoDB.NewDeviceDB(db))
	item, err := service.GetDevice(vars["id"])
	if err != nil {
		if err.Error() == "server error" {
			createError(w, "internal server error", http.StatusInternalServerError)
		} else {
			createError(w, "device not found", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(item)
	_, _ = w.Write(result)
}
