
build: build-server build-client

build-server:
	@go build -o bin/cmd-server cmd/server/server.go cmd/server/commands.go

build-client:
	@go build -o bin/cmd-client cmd/client/client.go