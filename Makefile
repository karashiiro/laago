SHELL := /bin/bash
.DEFAULT_GOAL := help 

help: ## Show this help
	@echo Dependencies: go
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build_dummy: ## Build the dummy application for testing
	-rm dummy.exe
	GOARCH=386 go build ./cmd/dummy/