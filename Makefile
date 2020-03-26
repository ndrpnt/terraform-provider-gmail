.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@awk 'BEGIN {FS = ":.*?## "} /^[0-9a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: fmt
fmt: ## Run go fmt against code
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code
	go vet ./...

.PHONY: build
build: ## Build provider binary
	GO111MODULE=on go build -o terraform-provider-gmail

.PHONY: test
test: fmt vet ## Run unit tests
	go test ./... $(TESTARGS)

.PHONY: testacc
testacc: fmt vet ## Run acceptance tests
	TF_ACC=1 go test ./... -timeout 20m -v $(TESTARGS)
