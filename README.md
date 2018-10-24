# Winter Is Coming - The game
_Protect the Wall_

## Background
I've chosen to implement an agnostic game engine that can be executed from different context, to comply with the exercise it is executed trough command line server/client architecture, but it could be plugged to a web server or behind a gRPC proxy server for example

## Libraries
I've deliberately avoided using libraries to maximize the chances of having to deal with Go most iconic types, structures and built-in packages; in a real world scenario I would spend some time to find what is best suited for the task at hands.

## Code style
Although there's inherent style in my code (like use NewX for constructors), I'm still flexible and as long as the team agrees on a standard I'm happy to follow it

## How to test
### Install the dependencies
The tests require a library to ease doing asserts and mocks  
Either get it globally or use the provide dependencies to install it as a local vendor  
```
go get github.com/stretchr/testify/assert
go get github.com/golang/mock
```
Using `dep` package manager  
```
dep ensure
```
### Run the tests
```
make tests
```
If GNU make file is not installed you can use:
```
go test -timeout 5s ./cmd/server ./game
```
### See code coverage
```
make coverage
```
If GNU make file is not installed you can use:
```
make tests-ci TEST_FLAGS=-coverprofile=coverage.out
go tool cover -html=coverage.out
```

## How to build  
**Disclaimer:** The version of Go used to develop this app is >= go1.10  
If you have installed GNU make you can invoke a helper command that will build it for you  
```
make build
```
Alternatively run the native command to build the server and the client
```
go build -o bin/cmd-server cmd/server/server.go cmd/server/commands.go
go build -o bin/cmd-client cmd/client/client.go
```

## How to run
Start the server in a terminal and the client in another, once finished playing you can kill them with Ctrl-c

You can pass an optional flag `address` in case the default port is in use
```
$ bin/cmd-server -address=:8122
Listening on address :8122

$ bin/cmd-client -address=127.0.0.1:8122
Connected to game server
```

### Build the server as docker image
Given you have `docker` and `make` installed in your machine you can also use it to build images for both the server and the client without requiring neither Go or `dep` installed

```
make docker-build
```
Then to start it
```
make docker-server
```
**NOTE:** Since no configuration is passed to the app it will read try to use the default port always

## Things I would improve
I've tried to stay true to the workday limit, I would love to implement some of the bonus points for the test, and also there's some other general things that I would improve:

* Parameterize the server and game options (port, board size, etc) with:
    - command line flags
    - environmental variables
    - config files 
* Add the the bonus points features
* Manage future dependencies with a package manager: dep, glide
* Improve logging
    - Use level thresholds and hide debug messages
    - Add context to debug problems (client, address)
* Handle OS signals
* Improve code coverage
* Use some library to ease the handling of the server/client communication
* Use a release tool to create cross-platform binaries and native packages, I've used [goreleaser](https://github.com/goreleaser/goreleaser) in the past in CI/CD pipelines with great results
* Add monitoring metrics
