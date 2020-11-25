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

start1:
	go build -o $(GOBIN)/$(PROJECTNAME)1 ./scraping-service/main.go || exit
	./bin/$(PROJECTNAME)1

start2:
	go build -o $(GOBIN)/$(PROJECTNAME)2 ./saving-service/main.go || exit
	./bin/$(PROJECTNAME)2

run:
	go run  ./main.go || exit
	./bin/$(PROJECTNAME)