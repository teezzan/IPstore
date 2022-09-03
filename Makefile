PROJECT_NAME:=IPstore

.PHONY: all
all: help

.PHONY: help
help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: mod
mod: ## Ensured Go package dependencies
	go mod tidy

.PHONY: test
test: ## Run unit tests
	go test -v ./...

.PHONY: benchmark
benchmark: ## Run benchmark 
	go test -run=XXX -bench=.