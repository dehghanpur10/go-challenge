package deviceRoute

import (
	"github.com/gorilla/mux"
	"go-challenge/controller/deviceController"
)

//SetDeviceRouter is func that set router of device model
func SetDeviceRouter(router *mux.Router) {
	router.HandleFunc("/devices", deviceController.GetDevice).Methods("GET")
	router.HandleFunc("/devices/{id}", deviceController.SetDevice).Methods("POST")
}
