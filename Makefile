APP_NAME=app
APP_VERSION=0.0.1
LDFLAGS=--ldflags '-X main.version=${APP_VERSION} -X main.appName=${APP_NAME} -extldflags "-static" -w -s'
NO_VENDOR=$(shell go list ./... | grep -v /vendor/)
OS=linux

.DEFAULT_GOAL := build

# Build app
build:
	CGO_ENABLED=0 GOOS=${OS} go build -v -a ${LDFLAGS} -tags netgo -installsuffix netgo \
    -o ./build/${APP_NAME} github.com/fabiorphp/echo-skeleton/cmd/app/
.PHONY: build

# Clean up
clean:
	@rm -fR ./build/ ./vendor/
.PHONY: clean

# Creates folders and download dependencies
configure:
	@mkdir -p ./build
	dep ensure -v
.PHONY: configure

# Run tests and generates html coverage file
cover:
	echo "mode: set" > ./build/coverage.out; \
    for i in ${NO_VENDOR}; do \
        go test -v -race -coverprofile=./build/cover.out $$i; \
        test -f ./build/cover.out && tail -n +2 ./build/cover.out >> ./build/coverage.out; \
    done; \
    go tool cover -html=./build/coverage.out -o ./build/coverage.html; \
    test -f ./build/cover.out && rm ./build/cover.out; \
    test -f ./build/coverage.out && rm ./build/coverage.out;
.PHONY: cover

# Install dependencies
depend:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install
	go get -u github.com/golang/dep/...
.PHONY: depend

# Format all go files
fmt:
	gofmt -s -w -l $(shell go list -f {{.Dir}} ./... | grep -v /vendor/)
.PHONY: fmt

# Run linters
lint:
	gometalinter.v2 --vendor --disable-all --enable=golint ./...
.PHONY: lint

# Run tests
test:
	go test -v -race ${NO_VENDOR}
.PHONY: test
