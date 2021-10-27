# jobbko
Highly concurrent Amazon SQS message scheduling service written in Go.

## Requirements
* [Go](https://golang.org)
* [Docker](https://www.docker.com)
* [AWS Command Line Interface](https://aws.amazon.com/cli/)

## Setup
```sh
$ make setup
```

This command will perform the following tasks:

* Build the `jobbko` binary
* Build the Docker image
* Spin up any dependency services
* Provision the required AWS resources

## Running
```sh
$ make run
```

If all has gone to plan, `jobbko` should now be running and accessible on `http://localhost:8000`.
