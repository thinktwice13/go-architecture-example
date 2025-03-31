.PHONY: setup test clean

setup:

test-all:
	@go test ./... -v

clean: