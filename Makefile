GOCMD=go
GOFMT=goreturns
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODOC=$(GOCMD) doc
GOGET=$(GOCMD) get
BINARY=myapp

all: fmt build test cover bench doc run
ci: build test cover bench
doc:
	$(GODOC) -all .
fmt:
	$(GOFMT) -l -w .
build:
	$(GOBUILD) -v -o $(BINARY)
buildlinux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v -o $(BINARY)
test:
	$(GOTEST) -race -v
bench:
	$(GOTEST) -parallel=4 -run="none" -benchtime="2s" -benchmem -bench=.
cover:
	$(GOTEST) -race -cover -covermode=atomic -coverprofile=coverage.out
run: build
	./$(BINARY)
