# Used by `image`, `push` & `deploy` targets, override as required
IMAGE_REG ?= ghcr.io
IMAGE_REPO ?= benc-uk/vuego-demoapp
IMAGE_TAG ?= latest

# Used by `deploy` target, sets Azure webap defaults, override as required
AZURE_RES_GROUP ?= temp-demoapps
AZURE_REGION ?= uksouth
AZURE_SITE_NAME ?= vuegoapp-$(shell git rev-parse --short HEAD)

# Used by `test-api` target
TEST_HOST ?= localhost:4000

# Don't change
SPA_DIR := spa
SRC_DIR := server
GOLINT_PATH := $(shell go env GOPATH)/bin/golangci-lint

.PHONY: help lint lint-fix image push run deploy undeploy clean test test-api test-report test-snapshot watch-server watch-spa .EXPORT_ALL_VARIABLES
.DEFAULT_GOAL := help

help:  ## 💬 This help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

lint: $(SPA_DIR)/node_modules  ## 🔎 Lint & format, will not fix but sets exit code on error 
	@$(GOLINT_PATH) > /dev/null || cd $(SRC_DIR); go get github.com/golangci/golangci-lint/cmd/golangci-lint
	cd $(SRC_DIR); $(GOLINT_PATH) run --modules-download-mode=mod *.go
	cd $(SPA_DIR); npm run lint

lint-fix: $(SPA_DIR)/node_modules  ## 📜 Lint & format, will try to fix errors and modify code
	@$(GOLINT_PATH) > /dev/null || cd $(SRC_DIR); go get github.com/golangci/golangci-lint/cmd/golangci-lint
	cd $(SRC_DIR); golangci-lint run --modules-download-mode=mod *.go --fix
	cd $(SPA_DIR); npm run lint-fix

image:  ## 🔨 Build container image from Dockerfile 
	docker build . --file build/Dockerfile \
	--tag $(IMAGE_REG)/$(IMAGE_REPO):$(IMAGE_TAG)

push:  ## 📤 Push container image to registry 
	docker push $(IMAGE_REG)/$(IMAGE_REPO):$(IMAGE_TAG)

run: $(SPA_DIR)/node_modules  ## 🏃 Run BOTH components locally using Vue CLI and Go server backend
	cd $(SRC_DIR); go run main.go routes.go &
	cd $(SPA_DIR); npm run serve

watch-server:  ## 👀 Run API server with hot reload file watcher, needs cosmtrek/air
	cd $(SRC_DIR); air

watch-spa: $(SPA_DIR)/node_modules  ## 👀 Run frontend SPA with hot reload file watcher
	cd $(SPA_DIR); npm run serve

deploy:  ## 🚀 Deploy to Azure Web App 
	az group create --resource-group $(AZURE_RES_GROUP) --location $(AZURE_REGION) -o table
	az deployment group create --template-file deploy/webapp.bicep \
		--resource-group $(AZURE_RES_GROUP) \
		--parameters webappName=$(AZURE_SITE_NAME) \
		--parameters webappImage=$(IMAGE_REG)/$(IMAGE_REPO):$(IMAGE_TAG) -o table 
	@echo "### 🚀 Web app deployed to https://$(AZURE_SITE_NAME).azurewebsites.net/"

undeploy:  ## 💀 Remove from Azure 
	@echo "### WARNING! Going to delete $(AZURE_RES_GROUP) 😲"
	az group delete -n $(AZURE_RES_GROUP) -o table --no-wait

test: $(SPA_DIR)/node_modules  ## 🎯 Unit tests for server and frontend 
	cd $(SRC_DIR); go test -v | tee server_tests.txt
	cd $(SPA_DIR); npm run test

test-report: test  ## 🎯 Unit tests for server and frontend (with report output)

test-snapshot:  ## 📷 Update snapshots for frontend tests
	cd $(SPA_DIR); npm run test-update

test-api: $(SPA_DIR)/node_modules .EXPORT_ALL_VARIABLES  ## 🚦 Run integration API tests, server must be running 
	$(SPA_DIR)/node_modules/.bin/newman run tests/postman_collection.json --env-var apphost=$(TEST_HOST)

clean:  ## 🧹 Clean up project
	rm -rf $(SPA_DIR)/dist
	rm -rf $(SPA_DIR)/node_modules
	rm -rf $(SRC_DIR)/server_tests.txt
	rm -rf $(SPA_DIR)/test*.html
	rm -rf $(SPA_DIR)/coverage

# ============================================================================

$(SPA_DIR)/node_modules: $(SPA_DIR)/package.json
	cd $(SPA_DIR); npm install --silent
	touch -m $(SPA_DIR)/node_modules

$(SPA_DIR)/package.json: 
	@echo "package.json was modified"
