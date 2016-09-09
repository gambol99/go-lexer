NAME=go-lexer
AUTHOR=gambol99
ROOT_DIR=${PWD}
GOVERSION=1.7.1
GIT_SHA=$(shell git --no-pager describe --always --dirty)
DEPS=$(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
PACKAGES=$(shell go list ./...)
LFLAGS ?= -X main.gitsha=${GIT_SHA}
VETARGS ?= -asmdecl -atomic -bool -buildtags -copylocks -methods -nilfunc -printf -rangeloops -shift -structtags -unsafeptr

.PHONY: test authors lint cover vet

default: test

golang:
	@echo "--> Go Version"
	@go version

docker-test:
	@echo "--> Compiling the project"
	${SUDO} docker run --rm -v ${ROOT_DIR}:/go/src/github.com/${AUTHOR}/${NAME} \
		-w /go/src/github.com/${AUTHOR}/${NAME} golang:${GOVERSION} make test

authors:
	@echo "--> Updating the AUTHORS"
	git log --format='%aN <%aE>' | sort -u > AUTHORS

deps:
	@echo "--> Installing build dependencies"
	@go get github.com/stretchr/testify/assert
	@go get github.com/davecgh/go-spew/spew

vet:
	@echo "--> Running go vet $(VETARGS) ."
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@go tool vet $(VETARGS) *.go

lint:
	@echo "--> Running golint"
	@which golint 2>/dev/null ; if [ $$? -eq 1 ]; then \
		go get -u github.com/golang/lint/golint; \
	fi
	@golint .

gofmt:
	@echo "--> Running gofmt check"
	@gofmt -s -l *.go \
	    | grep -q \.go ; if [ $$? -eq 0 ]; then \
            echo "You need to run the make format, we have file unformatted"; \
            gofmt -s -l *.go; \
						exit 1; \
	    fi

format:
	@echo "--> Running go fmt"
	@gofmt -s -w *.go

bench:
	@echo "--> Running go bench"
	@go test -v -bench=.

coverage:
	@echo "--> Running go coverage"
	@go test -coverprofile cover.out
	@go tool cover -html=cover.out -o cover.html

cover:
	@echo "--> Running go cover"
	@go test --cover

coveralls:
	@echo "--> Submitting to Coveralls"
	@go get github.com/mattn/goveralls

all: deps
	@echo "--> Running all the tests"
	@$(MAKE) test
	@$(MAKE) gofmt
	@$(MAKE) vet
	@$(MAKE) cover

test:
	@echo "--> Running the tests"
	@go test -v
	@$(MAKE) cover
