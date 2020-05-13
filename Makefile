.PHONY: clean run-dynamodb build run-local deploy

clean: 
	rm -rf ./bin
	
run-dynamodb:
	docker run -p 8000:8000 -v `pwd`/dbstore:/dbstore amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb -dbPath /dbstore

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/api lambda/main.go

run-local:
	go run cmd/main.go

deploy-dev: clean build
	sls deploy

