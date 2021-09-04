package deviceController

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go-challenge/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	input := models.Device{
		Id:          "333",
		DeviceModel: "a",
		Name:        "a",
		Note:        "a",
		Serial:      "a",
	}
	_ = os.Setenv("AWS_REGION", "us-west-2")
	_ = os.Setenv("TABLE_NAME", "Devices")

	CreateItem(t, input)
	_ = os.Unsetenv("AWS_REGION")
	_ = os.Unsetenv("TABLE_NAME")
	tests := []struct {
		name   string
		id     string
		status int
		output interface{}
	}{
		{name: "server error", status: 500, output: models.Error{
			Message: "internal server error",
		}, id: "01001010100101"},
		{name: "not found", status: 404, output: models.Error{
			Message: "device not found",
		}, id: "0101010105010fd"},
		{name: "ok", status: 200, output: input, id: input.Id},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name != "server error" {
				_ = os.Setenv("AWS_REGION", "us-west-2")
				_ = os.Setenv("TABLE_NAME", "Devices")
			}
			router := mux.NewRouter()
			router.HandleFunc("/devices/{id}", GetDevice).Methods("GET")

			req, _ := http.NewRequest(http.MethodGet, "/devices/"+test.id, nil)

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			assert.Equal(t, test.status, rr.Code)

			if test.status == 200 {
				var device models.Device
				_ = json.Unmarshal(rr.Body.Bytes(), &device)
				assert.Equal(t, test.output.(models.Device), device)
			} else {
				var message models.Error
				_ = json.Unmarshal(rr.Body.Bytes(), &message)
				assert.Equal(t, test.output.(models.Error), message)

			}
		})
	}

	DeleteItem(t, input.Id)
}
