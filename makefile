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
FRONT_DIR := frontend
SERVER_DIR := server
REPO_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
GOLINT_PATH := $(REPO_DIR)/bin/golangci-lint

.PHONY: help lint lint-fix image push run deploy undeploy clean test test-api test-report test-snapshot watch-server watch-spa .EXPORT_ALL_VARIABLES
.DEFAULT_GOAL := help

help:  ## 💬 This help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

lint: $(FRONT_DIR)/node_modules  ## 🔎 Lint & format, will not fix but sets exit code on error 
	@$(GOLINT_PATH) > /dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh
	cd $(SERVER_DIR); $(GOLINT_PATH) run --modules-download-mode=mod ./...
	cd $(FRONT_DIR); npm run lint

lint-fix: $(FRONT_DIR)/node_modules  ## 📜 Lint & format, will try to fix errors and modify code
	@$(GOLINT_PATH) > /dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh
	cd $(SERVER_DIR); golangci-lint run --modules-download-mode=mod *.go --fix
	cd $(FRONT_DIR); npm run lint-fix

image:  ## 🔨 Build container image from Dockerfile 
	docker build . --file build/Dockerfile \
	--tag $(IMAGE_REG)/$(IMAGE_REPO):$(IMAGE_TAG)

push:  ## 📤 Push container image to registry 
	docker push $(IMAGE_REG)/$(IMAGE_REPO):$(IMAGE_TAG)

run: $(FRONT_DIR)/node_modules  ## 🏃 Run BOTH components locally using Vue CLI and Go server backend
	cd $(SERVER_DIR); go run ./cmd &
	cd $(FRONT_DIR); npm run serve

watch-server:  ## 👀 Run API server with hot reload file watcher, needs cosmtrek/air
	cd $(SERVER_DIR); air -c .air.toml

watch-frontend: $(FRONT_DIR)/node_modules  ## 👀 Run frontend SPA with hot reload file watcher
	cd $(FRONT_DIR); npm run serve

build-frontend: $(FRONT_DIR)/node_modules  ## 🧰 Build and bundle the frontend SPA into dist
	cd $(FRONT_DIR); npm run build

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

test: $(FRONT_DIR)/node_modules  ## 🎯 Unit tests for server and frontend 
	cd $(SERVER_DIR); go test -v ./...
	cd $(FRONT_DIR); npm run test

test-snapshot:  ## 📷 Update snapshots for frontend tests
	cd $(FRONT_DIR); npm run test-update

test-api: $(FRONT_DIR)/node_modules .EXPORT_ALL_VARIABLES  ## 🚦 Run integration API tests, server must be running 
	$(FRONT_DIR)/node_modules/.bin/newman run tests/postman_collection.json --env-var BASE_URL=$(TEST_HOST)

clean:  ## 🧹 Clean up project
	rm -rf $(FRONT_DIR)/dist
	rm -rf $(FRONT_DIR)/node_modules
	rm -rf $(SERVER_DIR)/server_tests.txt
	rm -rf $(FRONT_DIR)/test*.html
	rm -rf $(FRONT_DIR)/coverage
	rm -rf $(REPO_DIR)/bin

# ============================================================================

$(FRONT_DIR)/node_modules: $(FRONT_DIR)/package.json
	cd $(FRONT_DIR); npm install --silent
	touch -m $(FRONT_DIR)/node_modules

$(FRONT_DIR)/package.json: 
	@echo "package.json was modified"
