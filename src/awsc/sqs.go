package awsc

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendEventToQueue(eventBody []byte, destination string) bool {
	queueUrl, getQueueError := getQueueUrlForDestination(destination)
	if getQueueError != nil {
		fmt.Println(getQueueError)
		return false
	}

	_, err := AwsClient.SqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(eventBody)),
		QueueUrl:    &queueUrl,
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func getQueueUrlForDestination(eventDestination string) (string, error) {
	result, err := AwsClient.SqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(eventDestination),
	})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return *result.QueueUrl, nil
}
