package awsc

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (c *AwsClient) InsertEvent(item ScheduledEventItem) {
	avs, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		c.logger.Errorf("Got error marshalling new item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      avs,
		TableName: aws.String("jobbko_scheduled_events"),
	}

	_, err2 := c.DynamoClient.PutItem(input)
	if err2 != nil {
		c.logger.Errorf("Got error calling PutItem: %s", err2)
	}
}

func (c *AwsClient) GetProcessableEvents(groupId int, limit int) []ScheduledEventItem {
	now := time.Now().Unix()

	result, err := c.DynamoClient.Query(&dynamodb.QueryInput{
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
		c.logger.Errorf("Error retrieving processable events from DynamoDB: %s", err)
	}

	var items []ScheduledEventItem
	marshErr := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	if marshErr != nil {
		c.logger.Errorf("Failed to unmarshal DynamoDB items: %s", marshErr)
	}

	return items
}

func (c *AwsClient) LockEvent(event ScheduledEventItem) bool {
	result, err := c.DynamoClient.UpdateItem(&dynamodb.UpdateItemInput{
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
		c.logger.Errorf("Failed to lock Event %s: %s", event.Id, err)
		return false
	}

	return *result.Attributes["state"].S == "LOCKED"
}

func (c *AwsClient) DeleteEvent(event ScheduledEventItem) {
	_, err := c.DynamoClient.DeleteItem(&dynamodb.DeleteItemInput{
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
		c.logger.Errorf("Failed to delete Event %s: %s", event.Id, err)
	}
}
