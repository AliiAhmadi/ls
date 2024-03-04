.DEFAULT_GOAL := help

help:
	@go build -o temp && ./temp -h && rm -rf temp

build:
	@go build -o ls

test: clear_cache
	@go test ./...

testv: clear_cache
	@go test -v ./...

clear_cache:
	@go clean -cache

.PHONY: help build test clear_cache testv