.PHONY: test

setup:
	@cd clean && go mod tidy
	@cd fcis && go mod tidy
	@cd hexagonal && go mod tidy
	@cd modular && go mod tidy
	@cd vertical && go mod tidy

test:
	@go test ./clean/... ./fcis/... ./hexagonal/... ./modular/... ./vertical/... -v