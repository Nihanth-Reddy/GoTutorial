package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"log"
)

type Person struct {
	UserId    string   `json:"user_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Age       int      `json:"age"`
	Courses   []string `json:"courses"`
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Create table Movies
	tableName := "GolangUsers"
	var userId string = "eba3b771-1479-4dab-8a8b-44f3143d83c2"

	item := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userId),
			},
		},
	}

	response, err := svc.GetItem(item)

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if response.Item == nil {
		log.Fatalf("Couldn't get the users with id: %s", userId)
	}

	fmt.Printf("Response Item returned is: %#v\n", response.Item)

	person := Person{}

	err = dynamodbattribute.UnmarshalMap(response.Item, &person)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %#v", err))
	}

	fmt.Printf("Found the person with user id -- %s:\n %#v", userId, person)
}
