  # Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    GODEPS=$(GOGET) -t -v ./...
    BINARY_NAME=CharCounterAPI

    all: test build
    build: 
			$(GOBUILD) -o $(BINARY_NAME) -v
    test: 
			$(GOTEST) -v ./...
    clean: 
			$(GOCLEAN)
			rm -f $(BINARY_NAME)
			rm -f $(BINARY_UNIX)
    run: 
			$(GOBUILD) -o $(BINARY_NAME) -v ./...
			./$(BINARY_NAME)
    deps: 
			$(GODEPS)

.PHONY: all test clean