package handlers

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	message := body
	resp := events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       fmt.Sprintf("%v", message),
		Headers: map[string]string{
			"Content-Type":                 "text/plain",
			"Access-Control-Allow-Origin":  "https://serverless-with-golang-frontend.vercel.app",
			"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type",
		},
	}

	return &resp, nil
}
