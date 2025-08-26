dev:
	go run main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./main.go
# 	go build -o bin/main main.go

start:
	go build -o bin/main main.go
	./bin/main

install:
	go install