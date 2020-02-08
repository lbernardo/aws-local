package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "My Database:" + os.Getenv("DATABASE") + " My table: " + os.Getenv("TABLE_NAME"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
