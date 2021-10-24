#!/bin/bash

############ Docker ############
echo "Starting necessary services..."
docker compose up -d
sleep 5
echo "Started necessary services"

############ AWS ############
echo "Provisioning required AWS resources..."

# DynamoDB
aws dynamodb create-table --table-name jobbko_scheduled_events \
      --attribute-definitions AttributeName=groupId,AttributeType=S AttributeName=id,AttributeType=S \
      --key-schema AttributeName=groupId,KeyType=HASH AttributeName=id,KeyType=RANGE \
      --billing-mode PAY_PER_REQUEST --endpoint-url http://localhost:4566 \
      >&1 > /dev/null

# SQS
aws --endpoint-url=http://localhost:4566 sqs create-queue \
      --queue-name jobbko_dispatched_events \
      >&1 > /dev/null

echo "Required AWS resources provisioned"

############ FIN ############
echo "Done!"
