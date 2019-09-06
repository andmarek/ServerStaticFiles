GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gogetem
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

build: go $(GOBUILD) -o $(BINARY_NAME) -v

run: go build blog.go testfile.go servetls.go 
	sudo ./$(BINARY_NAME)
