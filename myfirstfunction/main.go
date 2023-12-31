package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func Handler(event Request) (Response, error) {
	log.Printf("EVENT OBJECT: %+v", event)

	msg := "Hello from Lambda!"

	if name := event.QueryStringParameters["name"]; name != "" {
		msg = fmt.Sprintf("Hello %s!", name)
	}

	resp := Response{
		StatusCode: http.StatusOK,
		Body:       msg,
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
