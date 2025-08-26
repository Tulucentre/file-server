dev:
	go run main.go

build:
	go build -o bin/main main.go

start:
	go build -o bin/main main.go
	./bin/main

install:
	go install