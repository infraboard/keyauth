PROJECT_NAME=keyauth
MAIN_FILE=main.go
PKG := "github.com/infraboard/$(PROJECT_NAME)"
MOD_DIR := $(shell go env GOMODCACHE)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

BUILD_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_COMMIT := ${shell git rev-parse HEAD}
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION := $(shell go version | grep -o  'go[0-9].[0-9].*')
VERSION_PATH := "${PKG}/version"

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
	@go build -a -o dist/${PROJECT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" ${MAIN_FILE}

linux: dep ## Build the binary file
	@GOOS=linux GOARCH=amd64 go build -a -o dist/${PROJECT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" ${MAIN_FILE}

run: # Run Develop server
	@go run $(MAIN_FILE) start

init: # Init Service
	@go run $(MAIN_FILE) init

clean: ## Remove previous build
	@rm -f dist/*

push: # push git to multi repo
	@git push -u gitee
	@git push -u origin

gen: # Init Service
	@protoc -I=. --go_out=. --go_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} common/types/*.proto
	@protoc-go-inject-tag -input=common/types/*.pb.go
	@protoc -I=. --go_out=. --go_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} app/*/pb/*.proto
	@protoc-go-inject-tag -input=app/application/*.pb.go
	@protoc-go-inject-tag -input=app/department/*.pb.go
	@protoc-go-inject-tag -input=app/domain/*.pb.go
	@protoc-go-inject-tag -input=app/endpoint/*.pb.go
	@protoc-go-inject-tag -input=app/mconf/*.pb.go
	@protoc-go-inject-tag -input=app/micro/*.pb.go
	@protoc-go-inject-tag -input=app/namespace/*.pb.go
	@protoc-go-inject-tag -input=app/permission/*.pb.go
	@protoc-go-inject-tag -input=app/policy/*.pb.go
	@protoc-go-inject-tag -input=app/role/*.pb.go
	@protoc-go-inject-tag -input=app/session/*.pb.go
	@protoc-go-inject-tag -input=app/tag/*.pb.go
	@protoc-go-inject-tag -input=app/token/*.pb.go
	@protoc-go-inject-tag -input=app/user/*.pb.go
	@protoc-go-inject-tag -input=app/verifycode/*.pb.go
	@go generate ./...

install: dep# Install depence go package
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/infraboard/mcube/cmd/mcube@v1.1.2
	@go install github.com/infraboard/mcube/cmd/protoc-gen-go-ext@v1.1.2
	@go install github.com/infraboard/mcube/cmd/protoc-gen-go-http@v1.1.2

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
