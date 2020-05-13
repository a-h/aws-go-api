# Serverless API example.

A more realistic Serverless API template.

```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── cmd                         <-- Used to run the code locally
├── db                          <-- DynamoDB database access code
├── handlers                    <-- HTTP handler code
├── lambda                      <-- Lambda entry point, used to run the code within AWS
├── log                         <-- Structured logging configuration
├── respond                     <-- HTTP traffic logging and standard API responses
└── serverless.yaml             <-- Serverless Framework configuration
```

## Key features

* Run as a Docker container by packaging up the output of the `cmd` directory, or run Serverless.
* Uses DynamoDBLocal for local development.
* Uses standard Go HTTP library - not tied in to a specific Serverless implementation.
* JSON logging is provided by the "log" package. JSON logging allows metric extraction and analysis via CloudWatch Log Insights and CloudWatch Metrics.
* Simple HTTP handler dependency injection.
* JSON logging of HTTP response times.

## Local development

Run DynamoDB local (via Docker), then run the `cmd` package.

* `make run-dynamodb` (this blocks the Shell, so run it in a seperate process).
* `make run-local`

## Deployment

`make deploy-dev` will build the Go code and deploy it to AWS.

## TODO

* Add realistic examples of structured logging for business metrics.
* Add examples of HTTP handler unit tests, not just database integration tests.
* Add examples of JWT authorization decoding (ideally using Cognito).
* Add example of CORS middleware configuration.
* Add X-Ray configuration to allow tracing of request times downstream.

