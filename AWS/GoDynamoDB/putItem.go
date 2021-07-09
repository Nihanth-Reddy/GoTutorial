package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

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
	userId := uuid.NewString()

	person := Person{UserId: userId, FirstName: "Nihanth", LastName: "Kottam", Email: "nihanthreddy.kottam@cawstudios.com", Age: 25, Courses: []string{"Golang", "Python", "Javascript"}}
	// person := Person{UserId: userId, FirstName: "Nihanth", LastName: "Kottam", Email: "nihanthreddy.kottam@cawstudios.com", Age: 25, Courses: []string{"Golang", "Python", "Javascript"}}

	attributValue, err := dynamodbattribute.MarshalMap(person)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	item := &dynamodb.PutItemInput{
		Item:      attributValue,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(item)

	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	fmt.Printf("Successfully added user %s with id %s to the database.", person.FirstName, user_id)

}
