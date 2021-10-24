package awsc

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sirupsen/logrus"
)

type AwsClient struct {
	logger       *logrus.Logger
	AwsSession   *session.Session
	SqsClient    *sqs.SQS
	DynamoClient *dynamodb.DynamoDB
}

func InitAwsClient(logger *logrus.Logger) *AwsClient {
	session, err := session.NewSession(&aws.Config{
		Region:   aws.String("eu-west-2"),
		Endpoint: aws.String("http://localhost:4566"),
	})

	if err != nil {
		log.Fatal(err)
	}

	return &AwsClient{
		logger:       logger,
		AwsSession:   session,
		SqsClient:    sqs.New(session),
		DynamoClient: dynamodb.New(session),
	}
}

type ScheduledEventItem struct {
	Id      string      `json:"id"`
	GroupId string      `json:"groupId"`
	State   string      `json:"state"`
	Data    interface{} `json:"data"`
}
