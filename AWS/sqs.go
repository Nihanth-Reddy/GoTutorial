package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

type GroupObject struct {
	GroupId      string   `json:"groupId"`
	Date         string   `json:"date"`
	TimezoneName string   `json:"timezoneName"`
	UserIds      []string `json:"userIds"`
}

func GetQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SendMsg(sess *session.Session, queueURL *string) error {
	svc := sqs.New(sess)
	g := GroupObject{"123", "21-02-2120", "new", []string{"Hi", "Hello"}}
	res_str, _ := json.Marshal(g)
	new_str := string(res_str)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(new_str),
		QueueUrl:    queueURL,
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	groups := []GroupObject{{"123", "21-02-2120", "new", []string{"Hi", "Hello"}}, {"456", "21-02-2120", "new", []string{"Hi", "Hello"}}, {"789", "21-02-2120", "new", []string{"Hi", "Hello"}}, {"101112", "21-02-2120", "new", []string{"Hi", "Hello"}}, {"131415", "21-02-2120", "new", []string{"Hi", "Hello"}}}

	queue := "nihanthTestQueue"
	group_strs := []string{}
	for _, group := range groups {
		jBytes, _ := json.Marshal(group)
		group_strs = append(group_strs, string(jBytes))

	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Get URL of queue
	result, err := GetQueueURL(sess, &queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	queueURL := result.QueueUrl
	fmt.Printf("Queue URL: %v\n", *queueURL)

	err = sendBatch(sess, queueURL, group_strs)
	fmt.Println("Sent messages to the queue")

	// Sending single Message to Queue
	// err = SendMsg(sess, queueURL)
	// if err != nil {
	// 	fmt.Println("Got an error sending the message:")
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Sent message to queue ")
}

func sendBatch(sess *session.Session, queueURL *string, groups []string) error {
	svc := sqs.New(sess)
	var entries []*sqs.SendMessageBatchRequestEntry

	for _, group := range groups {
		var item sqs.SendMessageBatchRequestEntry
		item.Id = aws.String(getUUID())
		item.MessageBody = aws.String(group)
		entries = append(entries, &item)
	}
	var chunk []*sqs.SendMessageBatchRequestEntry
	count := 1
	for i, entry := range entries {
		chunk = append(chunk, entry)
		if count == 10 || (i == len(entries)-1) {
			res, err := svc.SendMessageBatch(&sqs.SendMessageBatchInput{
				Entries:  chunk,
				QueueUrl: queueURL,
			})
			fmt.Println(res)
			fmt.Println(err)
			count = 1
			chunk = nil
		} else {
			count += 1
		}
	}
	return nil
}

func getUUID() string {
	return uuid.NewString()
}
