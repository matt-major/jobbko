setup: build docker.build env.setup
cleanup: env.destroy
run: docker.run

build:
	@echo Building Jobbko binary...
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jobbko cmd/jobbko.go
	@echo Done

docker.build: build
	@echo Buidling Docker image...
	docker build . -t jobbko
	@echo Done

docker.run:
	@echo Running jobbko...
	docker run jobbko

env.setup:
	@echo Setting up environment...
	docker compose up -d

	@echo Waiting for resources to come online...
	sleep 10
	
	aws dynamodb create-table --table-name jobbko_scheduled_events \
      --attribute-definitions AttributeName=groupId,AttributeType=S AttributeName=id,AttributeType=S \
      --key-schema AttributeName=groupId,KeyType=HASH AttributeName=id,KeyType=RANGE \
      --billing-mode PAY_PER_REQUEST --endpoint-url http://localhost:4566 \
      >&1 > /dev/null

	aws --endpoint-url=http://localhost:4566 sqs create-queue \
      --queue-name jobbko_dispatched_events \
      >&1 > /dev/null

	@echo Done

env.destroy:
	@echo Tearing down environment...

	aws dynamodb delete-table \
      --table-name jobbko_scheduled_events \
      --endpoint-url http://localhost:4566 \
      >&1 > /dev/null

	aws --endpoint-url=http://localhost:4566 sqs delete-queue \
      --queue-url http://localhost:4566/000000000000/jobbko_dispatched_events \
      >&1 > /dev/null

	docker compose down

	@echo Done
