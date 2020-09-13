.PHONY: default

default: server cli

server:
	go build -o bin/server github.com/spilliams/blaseball/cmd/server

cli:
	go build -o bin/blase github.com/spilliams/blaseball/cmd/cli
