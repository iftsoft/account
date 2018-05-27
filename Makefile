GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOGET=$(GOCMD) get
GOFILE_NAME=main.go
BINARY_NAME=account

all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GORUN) $(GOFILE_NAME)
#deps:
#	$(GOGET) github.com/markbates/goth
