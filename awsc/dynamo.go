package awsc

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func InsertEvent(item ScheduledEventItem) {
	avs, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new item: %s", err)
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

func GetProcessableEvents(groupId int, limit int) []ScheduledEventItem {
	svc := dynamodb.New(AwsSession)

	now := time.Now().Unix()

	result, err := svc.Query(&dynamodb.QueryInput{
		TableName:              aws.String("jobbko_scheduled_events"),
		FilterExpression:       aws.String("#state = :state"),
		KeyConditionExpression: aws.String("#groupId = :groupId and #id < :id"),
		ExpressionAttributeNames: map[string]*string{
			"#groupId": aws.String("groupId"),
			"#state":   aws.String("state"),
			"#id":      aws.String("id"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":groupId": {
				S: aws.String(strconv.Itoa(groupId)),
			},
			":id": {
				S: aws.String(strconv.Itoa(int(now))),
			},
			":state": {
				S: aws.String("PENDING"),
			},
		},
		Limit: aws.Int64(int64(limit)),
	})

	if err != nil {
		fmt.Println(err)
	}

	var items []ScheduledEventItem
	marshErr := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	if marshErr != nil {
		fmt.Println(marshErr)
	}

	return items
}

func LockEvent(event ScheduledEventItem) bool {
	svc := dynamodb.New(AwsSession)

	result, err := svc.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String("jobbko_scheduled_events"),
		ExpressionAttributeNames: map[string]*string{
			"#state": aws.String("state"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":oldState": {
				S: aws.String("PENDING"),
			},
			":newState": {
				S: aws.String("LOCKED"),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"groupId": {
				S: aws.String(event.GroupId),
			},
			"id": {
				S: aws.String(event.Id),
			},
		},
		ConditionExpression: aws.String("#state = :oldState"),
		UpdateExpression:    aws.String("SET #state = :newState"),
		ReturnValues:        aws.String("ALL_NEW"),
	})

	if err != nil {
		fmt.Println("Failed to lock Event", event.Id, err)
		return false
	}

	return *result.Attributes["state"].S == "LOCKED"
}

func DeleteEvent(event ScheduledEventItem) {
	svc := dynamodb.New(AwsSession)

	_, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("jobbko_scheduled_events"),
		Key: map[string]*dynamodb.AttributeValue{
			"groupId": {
				S: aws.String(event.GroupId),
			},
			"id": {
				S: aws.String(event.Id),
			},
		},
	})

	if err != nil {
		fmt.Println("Failed to delete Event", event.Id, err)
	}
}
