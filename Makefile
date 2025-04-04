.PHONY: test

setup:
	@cd clean && go mod tidy
	@cd hexagonal && go mod tidy
	@cd modular && go mod tidy
	@cd vertical && go mod tidy

test:
	@go test ./clean/... ./hexagonal/... ./modular/... ./vertical/... -v

curl-test:
	@curl -iF "file=@sample.pdf" http://localhost:8080/documents