VERSION?=$(shell git describe --tags --abbrev=0)
REGISTRY?=docker.io/fangangan
IMAGE:=flex-gpu-device-plugin:$(VERSION)
DEV_IMAGE:=flex-gpu-device-plugin:dev

build:
	CGO_LDFLAGS_ALLOW='-Wl,--unresolved-symbols=ignore-in-object-files' go build -ldflags="-s -w" ./cmd/flex-gpu-device-plugin

image:
	docker build --no-cache -t $(IMAGE) .

push: image
	docker tag $(IMAGE) $(REGISTRY)/$(IMAGE)
	docker push $(REGISTRY)/$(IMAGE)

push.dev: image
	docker tag $(IMAGE) $(REGISTRY)/$(DEV_IMAGE)
	docker push $(REGISTRY)/$(DEV_IMAGE)
