package awsc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (c *AwsClient) SendEventToQueue(eventBody []byte, destination string) bool {
	queueUrl, getQueueError := c.getQueueUrlForDestination(destination)
	if getQueueError != nil {
		c.logger.Errorf("Failed to get Queue URL for %s: %s", destination, getQueueError)
		return false
	}

	_, err := c.SqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(eventBody)),
		QueueUrl:    &queueUrl,
	})

	if err != nil {
		c.logger.Errorf("Failed to send Event to Queue: %q", err)
		return false
	}

	return true
}

func (c *AwsClient) getQueueUrlForDestination(eventDestination string) (string, error) {
	result, err := c.SqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(eventDestination),
	})

	if err != nil {
		c.logger.Errorf("Failed to get Queue URL for %s: %s", eventDestination, err)
		return "", err
	}

	return *result.QueueUrl, nil
}
