package main

import (
	"github.com/gorilla/mux"
	"go-challenge/routes/deviceRoute"
	"log"
	"net/http"
)

func main(){
	router := mux.NewRouter()

	deviceRoute.SetDeviceRouter(router)
	http.Handle("/", router)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println(err)
	}
}
