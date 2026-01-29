.PHONY: hello_world
hello_world:
	@go version
	@cd 001_hello_world && go test

.PHONY: lint
lint:
	@which golangci-lint
	@golangci-lint run
