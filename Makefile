.DEFAULT_GOAL := help

APP?=project-ci
COMMIT_HASH=$(shell git rev-parse --short HEAD)
VERSION?=$(shell git rev-parse --short HEAD)
REGISTRY?=neohsudroid


.PHONY: setup
setup: ## setup go modules
	go mod tidy

.PHONY: clean
clean: ## cleans the binary
	go clean

.PHONY: run
run: ## runs go run the application
	go run -race cmd/app/app.go

.PHONY: build
build: clean setup ## build the application
	go build -ldflags="-w -s -X 'main.CommitHash=${COMMIT_HASH}' -X 'main.BuildTime=$(shell date)'" -o output/${APP} cmd/app/app.go

.PHONY: docker-build
docker-build: clean setup ## build docker image
	docker build --build-arg CommitHash=${COMMIT_HASH} --build-arg BuildTime="$(shell date)" --no-cache -t ${REGISTRY}/${APP}:${VERSION} -f build/Dockerfile .

.PHONY: docker-push
docker-push: docker-build ## push the docker image to registry
	docker push ${REGISTRY}/${APP}:${COMMIT_SHA}

.PHONY: help 
help: ## Prints this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'