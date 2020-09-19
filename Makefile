PROJECT_NAME := "dsa"

.PHONY: test

test:
	@echo 'Run all test files with race detection enabled'
	@go test -v -race ./...
