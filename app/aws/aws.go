package awsclient

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/matt-major/jobbko/app/scheduler"
)

var AwsSession *session.Session

func InitAwsSession() error {
	session, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:4566"),
	})

	if err != nil {
		return err
	}

	AwsSession = session

	return nil
}

func InsertEvent(event scheduler.ScheduledEvent) {
	avs, err := dynamodbattribute.MarshalMap(event)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      avs,
		TableName: aws.String("scheduled_events"),
	}

	svc := dynamodb.New(AwsSession)
	_, err2 := svc.PutItem(input)
	if err2 != nil {
		log.Fatalf("Got error calling PutItem: %s", err2)
	}
}
