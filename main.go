package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func(event any) (events.APIGatewayProxyResponse, error) {
		fmt.Println(event)

		resp := events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "Hello from Lambda!",
		}

		return resp, nil
	})
}
