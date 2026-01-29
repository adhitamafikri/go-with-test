.PHONY: lint
lint:
	@which golangci-lint
	@golangci-lint run

.PHONY: hello_world
hello_world:
	@go version
	@cd 001_hello_world && go test -v

.PHONY: integers
integers:
	@go version
	@cd 002_integers && go test -v

.PHONY: iterations
iterations:
	@go version
	@cd 003_iterations && go test -v -bench=. -benchmem
