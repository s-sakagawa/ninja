IMAGE_NAME := ninja/totp
CONTAINER_NAME := totp
WORKINGDIR := /go/src/
PWD := $(shell pwd)
VOLUME_DIR := /qrcode/

.PHONY: build
build:
		@docker build . -t $(IMAGE_NAME)

.PHONY: run
run:
		@docker run \
				--rm -it \
				--name $(CONTAINER_NAME) \
				--volume $(PWD)/$(VOLUME_DIR):$(WORKINGDIR)/$(VOLUME_DIR) \
				$(IMAGE_NAME)
