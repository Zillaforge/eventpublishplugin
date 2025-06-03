OWNER ?= ociscloud
PROJECT ?= EventPublishPlugin
ABBR ?= eventpublishplugin
IMAGE_NAME ?= event-publish-plugin
GOVERSION ?= 1.22.4
OS ?= ubuntu
ARCH ?= $(shell uname -m | sed 's/x86_64/amd64/' | sed 's/aarch64/arm64/')
PREVERSION ?= 0.1.2
VERSION ?= $(shell cat VERSION)
PWD := $(shell pwd)
GO_PROXY ?= "https://proxy.golang.org,http://proxy.pegasus-cloud.com:8078"

# Release Mode could be dev or prod,
# dev: default, will add commit id to version
# prod: will use version only
RELEASE_MODE ?= dev
COMMIT_ID ?= $(shell git rev-parse --short=8 HEAD)

sed = sed
ifeq ("$(shell uname -s)", "Darwin")	# BSD sed, like MacOS
	sed += -i ''
else	# GNU sed, like LinuxOS
	sed += -i''
endif

ifeq ($(RELEASE_MODE), prod)
    RELEASE_VERSION := $(VERSION)
else
    RELEASE_VERSION := $(VERSION)-$(COMMIT_ID)
endif

.PHONY: go-build
go-build:
	@echo "Build Binary"
	@go build -ldflags="-s -w" -o tmp/$(PROJECT)_$(RELEASE_VERSION)

.PHONY: build
build: go-build
ifeq ($(OS), ubuntu)
	@sh build/build-debian.sh
else
	@sh build/build-rpm.sh
endif

.PHONY: set-version
set-version:
	@echo "Set Version"
	@$(sed) -e'/$(PREVERSION)/{s//$(VERSION)/;:b' -e'n;bb' -e\} $(PWD)/constants/common.go
	@$(sed) -e'/$(PREVERSION)/{s//$(VERSION)/;:b' -e'n;bb' -e\} $(PWD)/etc/eventpublishplugin.yaml
	@$(sed) -e'/$(PREVERSION)/{s//$(VERSION)/;:b' -e'n;bb' -e\} $(PWD)/Makefile

.PHONY: build-container
build-container:
	@echo "Build Container"
	@go build -o build/busybox_image/tmp/eventpublishplugin
	@cp etc/eventpublishplugin.yaml build/busybox_image/tmp/eventpublishplugin.yaml

.PHONY: release
release: 
	@make set-version
	@make start-dev-persistent
	@rm -f eventpublishplugin
	@docker run --rm --name build-env -e GOPROXY=$(GO_PROXY) -e GOSUMDB="off" --network=host -v $(PWD):/home/eventpublishplugin -w /home/eventpublishplugin $(OWNER)/golang:$(GOVERSION)-$(OS)-$(ARCH) make OS=$(OS) build
	@docker run --rm --name build-env -e GOPROXY=$(GO_PROXY) -e GOSUMDB="off" --network=host -v $(PWD):/home/eventpublishplugin -v pegasus-cloud-eventpublishplugin:/tmp -w /home/eventpublishplugin $(OWNER)/golang:$(GOVERSION)-$(OS)-$(ARCH) cp -f etc/eventpublishplugin.yaml eventpublishplugin /tmp

.PHONY: release-image
release-image: 
	@make set-version
	@mkdir -p build/busybox_image/tmp
	@rm -rf tmp/container
	@docker run --name build-env --rm -e GOPROXY=$(GO_PROXY) -e GOSUMDB="off" --network=host -v $(PWD):/home/eventpublishplugin -w /home/eventpublishplugin $(OWNER)/golang:$(GOVERSION)-$(OS)-$(ARCH) make build-container
	@docker rm -f build-env
	@mkdir -p tmp/container
	@docker rmi -f $(OWNER)/$(IMAGE_NAME):$(RELEASE_VERSION)
	@docker build -t $(OWNER)/$(IMAGE_NAME):$(RELEASE_VERSION) build/busybox_image/
	@docker save $(OWNER)/$(IMAGE_NAME):$(RELEASE_VERSION) > tmp/container/$(ABBR)_$(RELEASE_VERSION).image.tar

.PHONY: release-image-file
release-image-file: release-image
	@rm -rf tmp/container
	@mkdir -p tmp/container
	@docker save $(OWNER)/$(IMAGE_NAME):$(RELEASE_VERSION) > tmp/container/$(ABBR)_$(RELEASE_VERSION).image.tar

.PHONY: push-image
push-image:
	@echo "Check Image $(OWNER)/$(IMAGE_NAME):$(RELEASE_VERSION)"
	@docker image inspect $(OWNER)/$(IMAGE_NAME):$(RELEASE_VERSION) --format="image existed"
	@echo "Push Image"
	@docker logout
	@echo "<DOCKER HUB KEY>" | docker login -u $(OWNER) --password-stdin
	@docker image push $(OWNER)/$(IMAGE_NAME):$(RELEASE_VERSION)
	@docker logout

.PHONY: start
start:
	@go env -w GOPROXY=$(GO_PROXY)
	@go env -w GOSUMDB="off"
	@go run main.go serve -c etc/eventpublishplugin.yaml --redis-channel eventpublishplugin

.PHONY: start-dev-env
start-dev-env:
	@make start-dev-persistent
	@make start-dev-system
	@make start-dev-service

.PHONY: start-dev-service
start-dev-service: docker-compose/service/docker-compose.*.yaml
	@for f in $^; do ARCH=$(ARCH) COMPOSE_IGNORE_ORPHANS=True docker-compose -f $${f} -p "pegasus-service" up -d --no-recreate || true; done

.PHONY: start-dev-system
start-dev-system: docker-compose/system/docker-compose.*.yaml
	@for f in $^; do COMPOSE_IGNORE_ORPHANS=True docker-compose -f $${f} -p "pegasus-system" up -d --no-recreate || true; done


.PHONY: start-dev-persistent
start-dev-persistent: docker-compose/persistent/docker-compose.*.yaml
	@for f in $^; do COMPOSE_IGNORE_ORPHANS=True docker-compose -f $${f} -p "pegasus-system" up -d --no-start || true; done

.PHONY: stop-dev-env # Stop and Remove current service only
stop-dev-env:
	COMPOSE_IGNORE_ORPHANS=True docker-compose -f docker-compose/service/docker-compose.${ABBR}.yaml -p "pegasus-service" down

.PHONY: stop-dev-all # Stop and Remove all dependency
stop-dev-all:
	@make stop-dev-service
	@make stop-dev-system

.PHONY: purge-dev-all # Stop and Remove all dependency include persistent network and volume
purge-dev-all:
	@make stop-dev-all
	@make clean-dev-persistent

.PHONY: stop-dev-service
stop-dev-service: docker-compose/service/docker-compose.*.yaml
	@for f in $^; do COMPOSE_IGNORE_ORPHANS=True docker-compose -f $${f} -p "pegasus-service" down -v; done

.PHONY: stop-dev-system
stop-dev-system: docker-compose/system/docker-compose.*.yaml
	@for f in $^; do COMPOSE_IGNORE_ORPHANS=True docker-compose -f $${f} -p "pegasus-system" down -v; done

.PHONY: clean-dev-persistent
clean-dev-persistent: docker-compose/persistent/docker-compose.*.yaml
	@for f in $^; do COMPOSE_IGNORE_ORPHANS=True docker-compose -f $${f} -p "pegasus-system" down -v; done
