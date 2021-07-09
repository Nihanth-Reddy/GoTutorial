package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	Message string `json:"message"`
}

func Greeting(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := req.QueryStringParameters["name"]
	res := Response{Message: "Hello " + name + "!, wlecome to go lambda"}

	response := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	response.StatusCode = 200

	response_str, _ := json.MarshalIndent(res, "", "  ")
	response.Body = string(response_str)
	fmt.Printf("Response is: %#v\n", response)
	return response, nil
}
