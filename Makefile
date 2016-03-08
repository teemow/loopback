PROJECT=loopback
ORGANIZATION=teemow

SOURCE := $(shell find . -name '*.go')
VERSION := $(shell cat VERSION)
PROJECT_PATH := $(GOPATH)/src/github.com/$(ORGANIZATION)

GOPATH := $(shell pwd)/.gobuild
GOVERSION := 1.6.0

ifndef GOOS
	GOOS := linux
endif
ifndef GOARCH
	GOARCH := amd64
endif

.PHONY: all clean run-tests deps bin install

all: deps $(PROJECT)

ci: clean all run-tests

clean:
	rm -rf $(GOPATH) $(PROJECT)

run-tests: 
	GOPATH=$(GOPATH) go test ./...

# deps
deps: .gobuild
.gobuild:
	mkdir -p $(PROJECT_PATH)
	cd $(PROJECT_PATH) && ln -s ../../../.. $(PROJECT)

	# Fetch private packages first (so `go get` skips them later)

	#
	# Fetch public packages
	GOPATH=$(GOPATH) go get -d github.com/$(ORGANIZATION)/$(PROJECT)

	#
	# Fetch test packages
	GOPATH=$(GOPATH) go get -d github.com/onsi/gomega
	GOPATH=$(GOPATH) go get -d github.com/onsi/ginkgo

# build
$(PROJECT): $(SOURCE) VERSION
	@echo Building for $(GOOS)/$(GOARCH)
	docker run \
	    --rm \
	    -v $(shell pwd):/usr/code \
	    -e GOPATH=/usr/code/.gobuild \
	    -e GOOS=$(GOOS) \
	    -e GOARCH=$(GOARCH) \
	    -w /usr/code \
	    golang:$(GOVERSION) \
	    go build -a -ldflags "-X main.projectVersion=$(VERSION) -X main.projectBuild=$(COMMIT)" -o $(PROJECT)

install: $(PROJECT)
	cp $(PROJECT) /usr/local/bin/

fmt:
	gofmt -l -w .
