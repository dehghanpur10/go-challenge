package deviceController

import (
	"encoding/json"
	"log"
	"net/http"
)

//GetDevice is controller for get device from dynamoDB
func GetDevice(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	marshal, err := json.Marshal(&struct {
		Name string `json:"name"`
	}{Name: "aaa"})
	write, err := w.Write(marshal)
	log.Println(write)
	if err != nil {
		return
	}

}
