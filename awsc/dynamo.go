package awsc

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func InsertEvent(event interface{}) {
	avs, err := dynamodbattribute.MarshalMap(event)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      avs,
		TableName: aws.String("jobbko_scheduled_events"),
	}

	svc := dynamodb.New(AwsSession)
	_, err2 := svc.PutItem(input)
	if err2 != nil {
		log.Fatalf("Got error calling PutItem: %s", err2)
	}
}