package deviceRoute

import (
	"github.com/gorilla/mux"
	"go-challenge/controller/deviceController"
)

//SetDeviceRouter is func that set router of device model
func SetDeviceRouter(router *mux.Router) {
	router.HandleFunc("/devices", deviceController.SetDevice).Methods("POST")
	router.HandleFunc("/devices/{id}", deviceController.GetDevice).Methods("GET")
}
