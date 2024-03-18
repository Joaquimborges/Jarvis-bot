.PHONY:run
run:
	@echo "### Start running the bot"
	@go run main.go

.PHONY:mocks
mocks:
	@echo "### Generating new mocks..."
	@go generate ./...

.PHONY: fmt
fmt:
	@echo "### Formatting the source code ###"
	@go fmt ./...

.PHONY: lint
lint:
	@echo "### Linting the source code ###"
	@golangci-lint run

.PHONY: vet
vet:
	@echo "### Checking for code issues ###"
	@go vet ./...

.PHONY: test
test:
	@echo "### Testing the app ###"
	@go test ./...
	@go clean --testcache