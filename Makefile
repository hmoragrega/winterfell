.PHONY: clean all

export TEST_FLAGS :=

deps:
	@dep ensure

tests:
	@make tests-ci TEST_FLAGS=-v

tests-ci:
	@go test -timeout 5s ${TEST_FLAGS} ./cmd/server ./game

coverage:
	@make tests-ci TEST_FLAGS=-coverprofile=coverage.out
	@go tool cover -html=coverage.out
	@rm coverage.out

include ops/makefiles/*.mk