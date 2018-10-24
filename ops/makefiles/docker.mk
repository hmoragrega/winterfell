export DOCKER_APP_PATH := /go/src/bitbucket.com/hmoragrega/winterfell
export DOCKER_GO_IMAGE := golang:1.11-alpine3.8
export DOCKER_SERVER_IMAGE := hmoragrega/winterfell-cmd-server
export DOCKER_CLIENT_IMAGE := hmoragrega/winterfell-cmd-client
export DOCKER_NETWORK := winterfell

docker-deps:
	@docker run -v `pwd`:${DOCKER_APP_PATH} ${DOCKER_GO_IMAGE} /bin/sh -c 'cd ${DOCKER_APP_PATH} \
	&& apk add --no-cache git curl \
	&& curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
	&& dep ensure -v'

docker-build:
	@docker build -t ${DOCKER_SERVER_IMAGE}:latest -f ops/docker/server/Dockerfile .

docker-server:
	@docker run --network=${DOCKER_NETWORK} -p 8100:8100 ${DOCKER_SERVER_IMAGE}:latest