# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=dc-dl
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v ./main.go
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Cross compilation
build-linux:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o $(BINARY_NAME) -v ./main.go
