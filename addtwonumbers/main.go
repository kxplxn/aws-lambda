package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handle(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	numAStr := event.QueryStringParameters["numA"]
	numBStr := event.QueryStringParameters["numB"]

	numA, err := strconv.Atoi(numAStr)
	if err != nil {
		log.Fatal(err)
	}

	numB, err := strconv.Atoi(numBStr)
	if err != nil {
		log.Fatal(err)
	}
	sum := numA + numB

	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprint(sum),
	}

	return resp, nil
}

func main() {
	lambda.Start(handle)
}
