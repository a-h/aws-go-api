package main

import (
	"context"
	"os"

	"github.com/a-h/aws-go-api/handlers"
	"github.com/a-h/aws-go-api/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	region := os.Getenv("REGION")
	tableName := os.Getenv("TABLE_NAME")
	log.Default.Info("start")
	h := handlers.All(handlers.Configuration{
		Region:           region,
		TableName:        tableName,
		UseDynamoDBLocal: false,
	})
	adapter := httpadapter.New(h)
	handler := func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return adapter.ProxyWithContext(ctx, req)
	}
	lambda.Start(handler)
}
