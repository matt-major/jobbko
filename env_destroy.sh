#!/bin/bash

############ AWS ############
echo "Removing AWS resources..."
# DynamoDB
aws dynamodb delete-table \
      --table-name jobbko_scheduled_events \
      --endpoint-url http://localhost:4566 \
      >&1 > /dev/null

# SQS
aws --endpoint-url=http://localhost:4566 sqs delete-queue \
      --queue-url http://localhost:4566/000000000000/jobbko_dispatched_events \
      >&1 > /dev/null

echo "Removed AWS resources"

############ Docker ############
echo "Removing Docker resources..."

docker compose down

echo "Removed Docker resources"

############ Fin ############
echo "Done!"