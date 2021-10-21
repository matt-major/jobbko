#!/bin/bash

# DynamoDB
aws dynamodb create-table --table-name jobbko_scheduled_events \
      --attribute-definitions AttributeName=groupId,AttributeType=S AttributeName=id,AttributeType=S \
      --key-schema AttributeName=groupId,KeyType=HASH AttributeName=id,KeyType=RANGE \
      --billing-mode PAY_PER_REQUEST --endpoint-url http://localhost:4566

# SQS
aws --endpoint-url=http://localhost:4566 sqs create-queue \
      --queue-name jobbko_dispatched_events
