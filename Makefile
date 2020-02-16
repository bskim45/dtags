NAME := dtags

EXTRACT_NUMBER := sed 's/[^0-9]*\([0-9]*\)[^0-9]*/\1/'
VERSION_FILE := common/current_version.go

VERSION_MAJOR := $(shell cat $(VERSION_FILE) | grep Major | $(EXTRACT_NUMBER))
VERSION_MINOR := $(shell cat $(VERSION_FILE) | grep Minor | $(EXTRACT_NUMBER))
VERSION_PATCH := $(shell cat $(VERSION_FILE) | grep Patch | $(EXTRACT_NUMBER))

VERSION ?= $(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_PATCH)
REVISION := $(shell git rev-parse --short HEAD)

DOCKER_IMAGE ?= bskim45/dtags

GOBASE := $(shell pwd)
GOLINT ?= bin/golangci-lint
GOLANG_DOCKER_IMAGE ?= golang:1.13.7

SRCS := $(shell find . -name '*.go' -type f)
PKGS := $(shell go list ./... | grep -v /vendor)

LDFLAGS := -ldflags="-s -w"

PLATFORMS := windows linux darwin freebsd
ARCHS := amd64 386

GOLANGCI_VERSION = 1.23.3

.DEFAULT_GOAL := bin/$(NAME)
bin/$(NAME): $(SRCS)
	@echo "> Building binary..."
	CGO_ENABLED=0 go build $(LDFLAGS) -o bin/$(NAME) -v

run: bin/$(NAME)
	bin/$(NAME)

.PHONY: clean
clean:
	go clean
	rm -rf bin/$(NAME)
	rm -rf dist/*

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint
bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ./bin/ v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

.PHONY: lint
lint: bin/golangci-lint
	$(GOLINT) run

.PHONY: fix
fix: bin/golangci-lint
	$(GOLINT) run --fix

.PHONY: bin/gox
bin/gox:
	go get github.com/mitchellh/gox

.PHONY: cross-build
cross-build: deps bin/gox
	rm -rf dist/*
	#for os in darwin linux windows; do \
#        		for arch in amd64 386; do \
#        			GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -a $(LDFLAGS) -o dist/$$os-$$arch/$(NAME); \
#        		done; \
#        	done
	gox -parallel 8 $(LDFLAGS) -os "$(PLATFORMS)" -arch "$(ARCHS)" \
		-osarch "linux/arm linux/arm64 freebsd/arm freebsd/arm64" \
		-output "dist/$(NAME)_{{.OS}}_{{.Arch}}/$(NAME)"

.PHONY: deps
deps:
	@echo "> Resolving missing dependencies..."
	go get $(get)

.PHONY: dist
dist: DIST_DIRS = find * -type d -exec
dist:
	@mkdir -p dist
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf $(NAME)-$(VERSION)-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r $(NAME)-$(VERSION)-{}.zip {} \; && \
	cd ..

.PHONY: release
release:
	git tag $(VERSION)
	git push origin $(VERSION)

.PHONY: test
test:
	#go test -cover -v $(PKGS)
	go test -cover -v ./...

.PHONY: check
check: test lint

install:
	go install

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: build-docker
build-docker:
	docker run --rm \
	    -v "$(GOBASE)":/go/src/$(NAME) \
	    -w /go/src/$(NAME) \
	    "$(GOLANG_DOCKER_IMAGE)" \
	    make all

.PHONY: docker-build-image
docker-build-image:
	docker build \
		-f Dockerfile \
		--build-arg VERSION=$(VERSION) \
		--build-arg VCS_REF=$(REVISON) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		-t $(DOCKER_IMAGE):$(VERSION) .

.PHONY: docker-tag
docker-tag:
	docker tag $(DOCKER_IMAGE):$(VERSION) $(DOCKER_IMAGE):latest
	docker tag $(DOCKER_IMAGE):$(VERSION) $(DOCKER_IMAGE):$(VERSION_MAJOR)

.PHONY: docker-publish
docker-publish:
	docker push $(DOCKER_IMAGE):$(VERSION)
	docker push $(DOCKER_IMAGE):$(VERSION_MAJOR)
	docker push $(DOCKER_IMAGE):latest

all: deps check bin/$(NAME)
