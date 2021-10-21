#!/bin/bash

# DynamoDB
aws dynamodb delete-table \
      --table-name jobbko_scheduled_events \
      --endpoint-url http://localhost:4566

# SQS
aws --endpoint-url=http://localhost:4566 sqs delete-queue \
      --queue-url http://localhost:4566/000000000000/jobbko_dispatched_events
