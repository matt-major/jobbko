package context

import (
	"github.com/matt-major/jobbko/src/awsc"
	"github.com/sirupsen/logrus"
)

type ApplicationContext struct {
	Logger    *logrus.Logger
	AwsClient *awsc.AwsClient
}

func CreateApplicationContext() *ApplicationContext {
	logger := logrus.New()
	awsClient := awsc.InitAwsClient(logger)

	context := ApplicationContext{
		Logger:    logger,
		AwsClient: awsClient,
	}

	context.Logger.Formatter = &logrus.TextFormatter{}

	return &context
}
