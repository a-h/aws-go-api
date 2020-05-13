package main

import (
	"net/http"
	"os"

	"github.com/a-h/aws-go-api/handlers"
	"github.com/a-h/aws-go-api/log"
	"go.uber.org/zap"
)

func main() {
	region := os.Getenv("REGION")
	tableName := os.Getenv("TABLE_NAME")
	log.Default.Info("start", zap.String("region", region), zap.String("table", tableName))
	h := handlers.All(handlers.Configuration{
		Region:           region,
		TableName:        tableName,
		UseDynamoDBLocal: true,
	})
	log.Default.Info("listening", zap.Int("port", 8080))
	http.ListenAndServe(":8080", h)
}
