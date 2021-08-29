package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillaMux"
	"github.com/gorilla/mux"
	_ "go-challenge/docs" // docs is generated by Swag CLI, you have to import it.
	"go-challenge/routes/deviceRoute"

	httpSwagger "github.com/swaggo/http-swagger"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	router := mux.NewRouter()
	deviceRoute.SetDeviceRouter(router)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)


	app := gorillamux.New(router)
	return app.Proxy(req)
}
// @title GO challenge
// @version 1.0
// @description Implement a simple Restful API on AWS
// @contact.name Mohammad Dehghanpour
// @contact.email m.dehghanpour10@gmail.com
// @host tvwvnaqfy9.execute-api.us-west-2.amazonaws.com/api
// @BasePath /
func main() {
	lambda.Start(Handler)
}

//env GOOS=linux go build -o bin/main main.go
