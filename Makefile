# Copyright 2017 The Caicloud Authors.
#
# The old school Makefile, following are required targets. The Makefile is written
# to allow building multiple binaries. You are free to add more targets or change
# existing implementations, as long as the semantics are preserved.
#
#   make        - default to 'build' target
#   make lint   - code analysis
#   make test   - run unit test (or plus integration test)
#   make build        - alias to build-local target
#   make build-local  - build local binary targets
#   make build-linux  - build linux binary targets
#   make container    - build containers
#   make push    - push containers
#   make clean   - clean up targets
#
# Not included but recommended targets:
#   make e2e-test
#
# The makefile is also responsible to populate project version information.
#
# Tweak the variables based on your project.
#

# Current version of the project.x`
VERSION ?= v0.9.2

# This repo's root import path (under GOPATH).
ROOT := github.com/caicloud/cyclone

# Target binaries. You can build multiple binaries for a single project.
TARGETS := server workflow/controller workflow/coordinator
IMAGES := server web workflow/controller workflow/coordinator resolver/git resolver/image resolver/kv

# Container image prefix and suffix added to targets.
# The final built images are:
#   $[REGISTRY]/$[IMAGE_PREFIX]$[TARGET]$[IMAGE_SUFFIX]:$[VERSION]
# $[REGISTRY] is an item from $[REGISTRIES], $[TARGET] is an item from $[TARGETS].
IMAGE_PREFIX ?= $(strip cyclone-)
IMAGE_SUFFIX ?= $(strip )

# Container registries.
REGISTRIES ?= test.caicloudprivatetest.com/release

#
# These variables should not need tweaking.
#

# A list of all packages.
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /test)

# Project main package location (can be multiple ones).
CMD_DIR := ./cmd

# Project output directory.
OUTPUT_DIR := ./bin

# Build direcotory.
BUILD_DIR := ./build

# Git commit sha.
COMMIT := $(shell git rev-parse --short HEAD)

# Golang standard bin directory.
BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

#
# Define all targets. At least the following commands are required:
#

# All targets.
.PHONY: lint test build container push

lint: $(GOMETALINTER)
	gometalinter ./pkg/... ./cmd/...

build: build-local

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

test:
	go test $(PKGS)

build-local:
	@for target in $(TARGETS); do                                                      \
	  CGO_ENABLED=0   GOOS=linux   GOARCH=amd64                                        \
	  go build -i -v -o $(OUTPUT_DIR)/$${target}                                       \
	    -ldflags "-s -w -X $(ROOT)/pkg/version.VERSION=$(VERSION)                      \
	              -X $(ROOT)/pkg/version.COMMIT=$(COMMIT)                              \
	              -X $(ROOT)/pkg/version.REPOROOT=$(ROOT)"                             \
	    $(CMD_DIR)/$${target};                                                         \
	done

build-linux:
	@for target in $(TARGETS); do                                                      \
	  for registry in $(REGISTRIES); do                                                \
	    docker run --rm                                                                \
	      -v $(PWD):/go/src/$(ROOT)                                                    \
	      -w /go/src/$(ROOT)                                                           \
	      -e GOOS=linux                                                                \
	      -e GOARCH=amd64                                                              \
	      -e GOPATH=/go                                                                \
	      -e CGO_ENABLED=0                                                             \
	        $${registry}/golang:1.10-alpine3.8                                         \
	          go build -i -v -o $(OUTPUT_DIR)/$${target}                               \
	            -ldflags "-s -w -X $(ROOT)/pkg/version.VERSION=$(VERSION)              \
	            -X $(ROOT)/pkg/version.COMMIT=$(COMMIT)                                \
	            -X $(ROOT)/pkg/version.REPOROOT=$(ROOT)"                               \
	            $(CMD_DIR)/$${target};                                                 \
	  done                                                                             \
	done

build-web:
	for registry in $(REGISTRIES); do                                                  \
	  docker run --rm                                                                  \
	    -v $(PWD)/web/:/app                                                            \
	    -w /app                                                                        \
	      $${registry}/node:8.9-alpine                                                 \
	        sh -c '                                                                    \
	          yarn;                                                                    \
	          yarn build';                                                             \
	done

build-web-local:
	sh -c '                                                                            \
	  cd web;                                                                          \
	  yarn;                                                                            \
	  yarn build'

container: build-linux
	@for image in $(IMAGES); do                                                        \
	  for registry in $(REGISTRIES); do                                                \
	    imageName=$(IMAGE_PREFIX)$${image/\//-}$(IMAGE_SUFFIX);                        \
	    docker build -t $${registry}/$${imageName}:$(VERSION)                          \
	      -f $(BUILD_DIR)/$${image}/Dockerfile .;                                      \
	  done                                                                             \
	done

container-local: build-local
	@for target in $(TARGETS); do                                                      \
	  for registry in $(REGISTRIES); do                                                \
	    image=$(IMAGE_PREFIX)$${target/\//-}$(IMAGE_SUFFIX);                           \
	    docker build -t $${registry}/$${image}:$(VERSION)                              \
	      -f $(BUILD_DIR)/$${target}/Dockerfile .;                                     \
	  done                                                                             \
	done

push: container
	@for target in $(TARGETS); do                                                      \
	  for registry in $(REGISTRIES); do                                                \
	    image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                                \
	    docker push $${registry}/$${image}:$(VERSION);                                 \
	  done                                                                             \
	done

gen: clean-generated
	bash tools/generator/autogenerate.sh

swagger:
	docker run --rm                                                                   \
	  -v $(PWD):/go/src/$(ROOT)                                                       \
	  -w /go/src/$(ROOT)                                                              \
	  -e GOOS=linux                                                                   \
	  -e GOARCH=amd64                                                                 \
	  -e GOPATH=/go                                                                   \
	  -e CGO_ENABLED=0                                                                \
	  $(REGISTRIES)/golang:1.10-alpine3.8                                             \
	  sh -c "apk add git &&                                                           \
	  go get -u github.com/caicloud/nirvana/cmd/nirvana &&                            \
	  go get -u github.com/golang/dep/cmd/dep &&                                      \
	  nirvana api --output web/public pkg/server/apis"

.PHONY: clean
clean:
	-rm -vrf ${OUTPUT_DIR}
clean-generated:
	-rm -rf ./pkg/k8s/informers
	-rm -rf ./pkg/k8s/clientset
	-rm -rf ./pkg/k8s/listers