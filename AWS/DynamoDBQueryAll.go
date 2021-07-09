package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const TableName = "GolangUsers"

type Person struct {
	UserId    string   `json:"user_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Age       int      `json:"age"`
	Courses   []string `json:"courses"`
}

func QueryAll() {
	dbClient := getDBClient()
	keyCond := expression.Key("user_id").Equal(expression.Value("eba3b771-1479-4dab-8a8b-44f3143d83c2"))
	// keyCond1 := expression.Key("user_id").Equal(expression.Value("08917254-7ed4-485d-a83b-bd84e1be2c0e"))
	fnameFilter := expression.Name("first_name").Equal(expression.Value("Nihanth"))
	lnameFilter := expression.Name("last_name").Equal(expression.Value("Kottam"))
	ageFilter := expression.Name("age").Equal(expression.Value(25))
	expr, _ := expression.NewBuilder().WithKeyCondition(keyCond).WithFilter(fnameFilter.And(lnameFilter).And(ageFilter)).Build()
	params := &dynamodb.QueryInput{
		TableName:                 aws.String(TableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}
	res, err := dbClient.Query(params)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf("The length of returned Items: %d\n", len(res.Items))
	var personIds []string
	for _, person := range res.Items {
		item := Person{}
		// log.Println(person)
		tStart := time.Now()
		err = dynamodbattribute.UnmarshalMap(person, &item)
		tEnd := tStart.Sub(time.Now())
		log.Println(tEnd)
		personIds = append(personIds, item.UserId)
	}
	// res_str, _ := json.Marshal(personIds)
	fmt.Println(personIds)

}

func getDBClient() dynamodbiface.DynamoDBAPI {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ServiceClient := dynamodb.New(sess)
	return ServiceClient
}

func main() {

	QueryAll()
}
