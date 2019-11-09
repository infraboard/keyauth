BINARY_NAME=keyauth
MAIN_FILE_PAHT=main.go

all: test build
run:
		@go build -o ${BINARY_NAME} ${MAIN_FILE_PAHT}
		@./${BINARY_NAME} service start
clean:
		@go clean .
		@rm -f ${BINARY_NAME}
test:
		go test -v ./...
build: local
linux:
		@sh ./build/build.sh linux ${BINARY_NAME} ${MAIN_FILE_PAHT}
local:
		@sh ./build/build.sh local ${BINARY_NAME} ${MAIN_FILE_PAHT}
docker:
		@sh ./build/build.sh docker ${BINARY_NAME} ${MAIN_FILE_PAHT}
image:
		@sh ./build/build.sh image ${BINARY_NAME} ${MAIN_FILE_PAHT} ${IMAGE_PREFIX}
