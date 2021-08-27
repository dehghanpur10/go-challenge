package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillaMux"
	"github.com/gorilla/mux"
	"go-challenge/routes/deviceRoute"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	router := mux.NewRouter()
	deviceRoute.SetDeviceRouter(router)

	app := gorillamux.New(router)
	return app.Proxy(req)
}
func main() {
	lambda.Start(Handler)
}

//env GOOS=linux go build -o bin/main main.go
