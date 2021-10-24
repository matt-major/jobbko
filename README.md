# jobbko
Highly concurrent Amazon SQS message scheduling service written in Go.

## Development

### Environment Setup

Run the following script to setup LocalStack and the required AWS resources:
* `./env_setup.sh`

Now, when you run `jobbko` it will use the `sqs` and `dynamo` resources provisioned by this setup script.

When you're finished, you can clear everything down with the following script:
* `./env_destroy.sh`
