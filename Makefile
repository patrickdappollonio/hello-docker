IMAGE_TAG ?= patrickdappollonio/hello-docker
BIN_NAME = hello-docker

default: build

build:
	go build -a -tags netgo -ldflags '-s -w' -o $$(pwd)/$(BIN_NAME)

clean:
	rm -rf $$(pwd)/$(BIN_NAME)

docker:
	docker build --pull=true --rm=true -t $(IMAGE_TAG) .

ci:
	$(MAKE) build
	$(MAKE) docker
	$(MAKE) clean

.PHONY: build clean docker ci

.NOTPARALLEL:
