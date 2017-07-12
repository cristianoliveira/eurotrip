.PHONY: all start build test

SOURCE=src/github.com/cristianoliveira/eurotrip
GOPATH=$(shell pwd)

all: test build

build:
	@GOPATH=$(GOPATH) go build -o bin/eurotrip "$(SOURCE)/main.go"

start:
	@GOPATH=$(GOPATH) go run "$(SOURCE)/main.go" $(GOPATH)/data/example

test:
	@GOPATH=$(GOPATH) go test ./... -v

test-load: build
	@echo "Running the loading test. Please ensure that server is running on 8088."

	@echo "Installing dependencies."
	@go get -u github.com/tsenart/vegeta

	@echo "Running... it may take some time."
	@echo "GET http://localhost:8088/api/direct?dep_sid=114&arr_sid=152" | vegeta attack -duration=10s | tee results.bin | vegeta report

test-bigfile: build
	@echo "Running test using a big file. Loading it may take some time (around 5 min) ..."
	$(GOPATH)/bin/eurotrip $(GOPATH)/data/example_big


watch:
	find * | entr make test
