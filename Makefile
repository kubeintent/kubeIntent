export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

build: # @HELP build the Go binaries and run all validations (default)
build: build-kube-dsl

build-kube-dsl: deps
	go build -o build/_output/kubedsl ./cmd/kubedsl

test: # @HELP run the unit tests and source code validation
test: build linters
	go test github.com/adibrastegarnia/kubeDSL/pkg/...
	go test github.com/adibrastegarnia/kubeDSL/cmd/...

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

all: build test

clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
