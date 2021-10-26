# jobbko
Highly concurrent Amazon SQS message scheduling service written in Go.

## Development

### Setup

Run the following script to setup LocalStack and the required AWS resources:
```sh
$ ./env_setup.sh
```

Now, when you run `jobbko` it will use the `sqs` and `dynamo` resources provisioned by this setup script.

When you're finished, you can clear everything down with the following script:
```sh
$ ./env_destroy.sh
```

### Build

```sh
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jobbko ./src
$ docker build . -t jobbko
```

## Running

```sh
$ docker run jobbko
```
