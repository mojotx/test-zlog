#@GO=env GOOS=linux GOARCH=amd64 GOPRIVATE=github.com/noiseaware GO111MODULE=on go
GO=env GOOS=darwin GOARCH=amd64 GOPRIVATE=github.com/noiseaware GO111MODULE=on go
GOBIN=$(HOME)/go/bin

TAGV=$(shell git describe --always --dirty )
PROJECT:=$(shell basename $(shell git rev-parse --show-toplevel ))
PROJECT_LC:=$(shell echo $(PROJECT) | tr '[:upper:]' '[:lower:]' )

MAIN=./main.go

.PHONY: all cache-clean check clean clean docker-clean fmt format lint mod-cache-clean regen test-cache-clean vet main

all: main

main:
	$(GO) build -o $@ $(MAIN)

squeaky-clean: clean cache-clean docker-clean

clean: cache-clean
	$(RM) main

cache-clean:
	$(GO) clean -x -cache

mod-cache-clean:
	$(GO) clean -x -modcache

test-cache-clean:
	$(GO) clean -x -testcache


docker-clean:
	docker ps -qa --no-trunc --filter "status=exited" --filter "status=created" | xargs docker rm
	docker images -qa --no-trunc --filter "dangling=true" | xargs docker image rm
	docker images -qa --no-trunc --filter "reference=$(REPO):$(TAGV)" | xargs docker image rm --force

fmt:
	go fmt ./...
	go mod edit -fmt

vet:
	@echo Go vet...
	$(GO) vet ./...
	@sh -c "echo '\x1B[1;30m'"
	which findcall || ( $(GO) get -v -u golang.org/x/tools/go/analysis/passes/findcall/cmd/findcall ; go mod tidy; )
	@sh -c "echo '\x1B[0m'"
	@sh -c "echo '\x1B[36mgo vet findcall\x1B[0m'"
	$(GO) vet -vettool=$(GOBIN)/findcall ./...
	@echo
	@sh -c "echo '\x1B[1;30m'"
	which lostcancel > /dev/null || ( $(GO) get -v -u golang.org/x/tools/go/analysis/passes/lostcancel/cmd/lostcancel ; go mod tidy; )
	@sh -c "echo '\x1B[0m'"
	@sh -c "echo '\x1B[36mgo vet lostcancel\x1B[0m'"
	$(GO) vet -vettool=$(GOBIN)/lostcancel ./...
	@echo
	@sh -c "echo '\x1B[1;30m'"
	which nilness > /dev/null || ( $(GO) get -v -u golang.org/x/tools/go/analysis/passes/nilness/cmd/nilness ; go mod tidy; )
	@sh -c "echo '\x1B[0m'"
	@sh -c "echo '\x1B[36mgo vet nilness\x1B[0m'"
	$(GO) vet -vettool=$(GOBIN)/nilness ./...
	@echo
	@sh -c "echo '\x1B[1;30m'"
	which shadow > /dev/null || ( $(GO) get -v -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow ; go mod tidy; )
	@sh -c "echo '\x1B[0m'"
	@sh -c "echo '\x1B[36mgo vet shadow\x1B[0m'"
	$(GO) vet -vettool=$(GOBIN)/shadow ./...
	@echo
	@sh -c "echo '\x1B[1;30m'"
	which stringintconv > /dev/null || ( $(GO) get -v -u golang.org/x/tools/go/analysis/passes/stringintconv/cmd/stringintconv ; go mod tidy; )
	@sh -c "echo '\x1B[0m'"
	@sh -c "echo '\x1B[36mgo vet stringintconv\x1B[0m'"
	$(GO) vet -vettool=$(GOBIN)/stringintconv ./...
	@echo
	@sh -c "echo '\x1B[1;30m'"
	which unmarshal > /dev/null || ( $(GO) get -v -u golang.org/x/tools/go/analysis/passes/unmarshal/cmd/unmarshal ; go mod tidy; )
	@sh -c "echo '\x1B[0m'"
	@sh -c "echo '\x1B[36mgo vet unmarshal\x1B[0m'"
	$(GO) vet -vettool=$(GOBIN)/unmarshal ./...

# alias
format: fmt

lint:
	golint $(MAIN)

check: fmt vet fix lint

fix:
	$(GO) fix ./...

debug: main
	env LOG_LEVEL=debug ./main
