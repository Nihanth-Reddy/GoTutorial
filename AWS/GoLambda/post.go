package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Course string `json:"course"`
}

type PostResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func main() {
	lambda.Start(handle_post)
}

func handle_post(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Incoming request is: %#v\n", req)
	switch req.HTTPMethod {
	case "GET":
		return Greeting(req)
	case "POST":
		var request Request
		err := json.Unmarshal([]byte(req.Body), &request)
		if err != nil {
			fmt.Printf("Recieved an error while unmarshalling the request body: %v\n", err)
		}
		response := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
		response.StatusCode = 200
		res_string, err := json.MarshalIndent(PostResponse{Status: true, Message: fmt.Sprintf("%s has successfully been enrolled to %s", request.Name, request.Course)}, "", "  ")
		if err != nil {
			fmt.Printf("Recieved an error while marshalling the response body: %v\n", err)
		}
		response.Body = string(res_string)
		fmt.Printf("Response is: %#v\n", response)
		return response, nil
	default:
		fmt.Printf("Recieved an unhandled status code: %v\n", req.HTTPMethod)
		response := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
		response.StatusCode = 405
		fmt.Printf("Response is: %#v\n", response)
		return response, nil
	}
}
