.DEFAULT_GOAL := build

install:
	@go get -u github.com/jstemmer/go-junit-report
.PHONY: install

init: install
	@go mod vendor
.PHONY: init

build: init
	@go build -a -mod=vendor -o app
.PHONY: build

fmt:
	@go fmt \
		golang-backend \
		golang-backend/controllers \
		golang-backend/consts
.PHONY: fmt

test:
	@go test -v  ./...
.PHONY: test

test-junit:
	@go test -v  ./... | go-junit-report > junit.xml
.PHONY: test-junit
