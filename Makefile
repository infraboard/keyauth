PROJECT_NAME=keyauth
MAIN_FILE=main.go
PKG := "github.com/infraboard/$(PROJECT_NAME)"
MOD_DIR := $(shell go env GOMODCACHE)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep lint vet test test-coverage build clean

all: build

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	@cat cover.out >> coverage.txt

build: dep ## Build the binary file
	@sh ./build/build.sh local dist/$(PROJECT_NAME) $(MAIN_FILE)

linux: dep ## Build the binary file
	@sh ./build/build.sh linux dist/$(PROJECT_NAME) $(MAIN_FILE)

run: # Run Develop server
	@go run $(MAIN_FILE) start

init: # Init Service
	@go run $(MAIN_FILE) init

clean: ## Remove previous build
	@rm -f dist/*

push: # push git to multi repo
	@git push -u gitee
	@git push -u origin

codegen: # Init Service
	@protoc -I=. -I${MOD_DIR} --go-ext_out=. --go-ext_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} common/types/*.proto
	@protoc -I=. -I${MOD_DIR} --go-ext_out=. --go-ext_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} --go-http_out=. --go-http_opt=module=${PKG} pkg/*/pb/*.proto
	@go generate ./...

install: dep# Install depence go package
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/infraboard/mcube/cmd/mcube@v1.1.2
	@go install github.com/infraboard/mcube/cmd/protoc-gen-go-ext@v1.1.2
	@go install github.com/infraboard/mcube/cmd/protoc-gen-go-http@v1.1.2

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
