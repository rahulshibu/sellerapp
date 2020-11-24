# Author Rahul V R <rahul.vr@accubits.com>
PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GO=$(shell which go)

install:
	go mod download

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -ldflags="-s -w" -o $(GOBIN)/$(PROJECTNAME) ./main.go
	chmod +x $(GOBIN)/$(PROJECTNAME) 

start:
	go build -o $(GOBIN)/$(PROJECTNAME) ./main.go || exit
	./bin/$(PROJECTNAME)

run:
	go run  ./main.go || exit
	./bin/$(PROJECTNAME)