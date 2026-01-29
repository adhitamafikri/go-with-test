.PHONY: exec
exec:
	@go version
	@go run ./cmd/main.go

.PHONY: lint
lint:
	@which golangci-lint
	@golangci-lint run
