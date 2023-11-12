SHELL = /bin/bash
CURRENT_DIRECTORY = $(shell pwd)

# Go variables
GOFILES = $(shell find . -type f -name '*.go' -not -path "*/mock/*.go" -not -path "*.pb.go" -not -path "*_eventgen.go" -not -path "*_gen.go")

.PHONY: all
all: dep generate ## Runs dep generate

.PHONY: dep
dep: ## Install dependencies
	@go install github.com/google/wire/cmd/wire@latest
	@go install entgo.io/ent/cmd/ent@latest
	@go install github.com/golang/mock/mockgen@latest
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	@go mod tidy
	@go mod download

.PHONY: generate
generate: ## Generate code
	@./api-definitions/events/scripts/generate-all.sh
	@GOFLAGS=-mod=mod go generate ./...

.PHONY: lint
lint: ## run golanci-lint locally
	@terraform fmt -check -diff -recursive
	@docker pull golangci/golangci-lint:latest-alpine
	@docker run --rm -v $(CURRENT_DIRECTORY):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run -v

.PHONY: format
format: ## Format the source
	@terraform fmt -diff -recursive
	@goimports -w $(GOFILES)

.PHONY: help
.DEFAULT_GOAL := help
help: ## Get help output
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
