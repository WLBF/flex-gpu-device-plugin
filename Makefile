VERSION?=$(shell git describe --tags --abbrev=0)
REGISTRY?=docker.io/fangangan
IMAGE:=gpu-share-device-plugin:$(VERSION)
DEV_IMAGE:=gpu-share-device-plugin:dev

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/gpu-share-device-plugin

image:
	docker build --no-cache -t $(IMAGE) .

push: image
	docker tag $(IMAGE) $(REGISTRY)/$(IMAGE)
	docker push $(REGISTRY)/$(IMAGE)

push.dev: image
	docker tag $(IMAGE) $(REGISTRY)/$(DEV_IMAGE)
	docker push $(REGISTRY)/$(DEV_IMAGE)
