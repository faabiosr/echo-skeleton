.PHONY: depend clean build build-docker test test-coverage

APP_NAME=myapp
APP_PATH=$(shell head -n 1 ./glide.yaml | awk '{print $$2}')
APP_VERSION=0.0.1

LDFLAGS=--ldflags '-X main.version=${APP_VERSION} -X main.appName=${APP_NAME} -extldflags "-static" -w'
OS=linux
PACKAGES=$(shell find ./handler -type d)

DOCKER_NS=mydockerns

.DEFAULT_GOAL := build

depend:
	@command -v glide > /dev/null 2>&1 || ( echo "Please install Glide https://github.com/Masterminds/glide" && exit 1 )
	@glide install

clean:
	@rm -fR vendor/ glide.lock ./bin ./.glide/ ${APP_NAME}

build: depend
ifeq ($(BUILD),docker)
	@command -v docker > /dev/null 2>&1 || ( echo "Please install Docker https://docs.docker.com/engine/installation/" && exit 1 )
	@docker run --rm \
        -v "$(shell pwd)":/go/src/${APP_PATH} \
        -w /go/src/${APP_PATH} \
        fabiorphp/golang-glide:1.8 sh -c "make OS=${OS} APP_NAME=${APP_NAME} APP_VERSION=${APP_VERSION}"
else
	CGO_ENABLED=0 GOOS=${OS} go build -a ${LDFLAGS} -tags netgo -installsuffix netgo -v -o ./bin/${APP_NAME}
endif

pack:
	@command -v docker > /dev/null 2>&1 || ( echo "Please install Docker https://docs.docker.com/engine/installation/" && exit 1 )
	@docker build -t ${DOCKER_NS}/${APP_NAME}:${APP_VERSION} --build-arg APP_NAME=${APP_NAME} -f ./Dockerfile .

test:
	@$(foreach pkg, $(PACKAGES),\
        go test $(pkg); \
    )

test-coverage:
	@echo "mode: set" > coverage.out
	@$(foreach pkg, $(PACKAGES),\
        go test -coverprofile=cover.out $(pkg); \
        tail -n +2 cover.out >> coverage.out; \
    )
	@go tool cover -html=coverage.out -o coverage.html
	@rm cover.out coverage.out
