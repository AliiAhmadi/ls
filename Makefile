.DEFAULT_GOAL := help

help:
	@go build -o temp && ./temp -h && rm -rf temp

build:
	@go build -o ls