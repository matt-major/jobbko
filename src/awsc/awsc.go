package awsc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var AwsClient *AwsClientStruct

type AwsClientStruct struct {
	AwsSession   *session.Session
	SqsClient    *sqs.SQS
	DynamoClient *dynamodb.DynamoDB
}

func InitAwsClient() error {
	session, err := session.NewSession(&aws.Config{
		Region:   aws.String("eu-west-2"),
		Endpoint: aws.String("http://localhost:4566"),
	})

	if err != nil {
		return err
	}

	AwsClient = &AwsClientStruct{
		AwsSession:   session,
		SqsClient:    sqs.New(session),
		DynamoClient: dynamodb.New(session),
	}

	return nil
}

type ScheduledEventItem struct {
	Id      string      `json:"id"`
	GroupId string      `json:"groupId"`
	State   string      `json:"state"`
	Data    interface{} `json:"data"`
}
