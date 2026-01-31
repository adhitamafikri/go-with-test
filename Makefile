.PHONY: lint
lint:
	@which golangci-lint
	@golangci-lint run

.PHONY: hello_world
hello_world:
	@go version
	@cd 001_hello_world && go test -v -cover

.PHONY: integers
integers:
	@go version
	@cd 002_integers && go test -v -cover

.PHONY: iterations
iterations:
	@go version
	@cd 003_iterations && go test -v -bench=. -benchmem -cover

.PHONY: arrays_slices
arrays_slices:
	@go version
	@cd 004_arrays_slices && go test -v -bench=. -benchmem -cover

.PHONY: structs_methods_interfaces
structs_methods_interfaces:
	@go version
	@cd 005_structs_methods_interfaces && go test -v -bench=. -benchmem -cover

.PHONY: pointers_and_errors
pointers_and_errors:
	@go version
	@cd 006_pointers_and_errors && go test -v -bench=. -benchmem -cover

.PHONY: maps
maps:
	@go version
	@cd 007_maps && go test -v -bench=. -benchmem -cover

.PHONY: dependency_injection
dependency_injection:
	@go version
	@cd 008_dependency_injection && go test -v -bench=. -benchmem -cover

.PHONY: mocking
mocking:
	@go version
	@cd 009_mocking && go test -v -bench=. -benchmem -cover
