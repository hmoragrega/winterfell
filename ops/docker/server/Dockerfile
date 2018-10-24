FROM golang:1.11-alpine3.8 AS build-env

COPY . /go/src/github.com/hmoragrega/winterfell

WORKDIR /go/src/github.com/hmoragrega/winterfell
RUN go build -o /bin/cmd-server cmd/server/server.go cmd/server/commands.go

# final stage
FROM alpine

COPY --from=build-env /bin/cmd-server /bin

CMD ["cmd-server"]