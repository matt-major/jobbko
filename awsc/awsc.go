package awsc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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
