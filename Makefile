BUILD_VERSION = v1.0.0
BINARY_NAME = gnm

DIST_PATH=./dist
LD_FLAGS = -ldflags "-X main.Version=$(BUILD_VERSION)"

## build: create binary in ./dist folder for your current platform. Use this option if you build it for personal use.
.PHONY: build
build:
	@rm -rf ./$(DIST_PATH)/*
	@echo 'Building'
	go build ${LD_FLAGS} -o $(DIST_PATH)/$(BINARY_NAME) ./cmd/gnm/*.go

## clean
.PHONY: clean
clean:
	@go clean
	@rm -rf ./$(DIST_PATH)/*

.PHONY: release
release:
	@rm -rf ./$(DIST_PATH)/*

	# Build for mac
	go build ${LD_FLAGS} -o $(DIST_PATH)/mac64-$(BINARY_NAME) ./cmd/gnm/*.go
	@tar czvf ./dist/${BINARY_NAME}-mac64-${BUILD_VERSION}.tar.gz ./dist/mac64-${BINARY_NAME}

	# Build for arm
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build ${LD_FLAGS} -o $(DIST_PATH)/arm64-$(BINARY_NAME) ./cmd/gnm/*.go
	@tar czvf ./dist/${BINARY_NAME}-arm64-${BUILD_VERSION}.tar.gz ./dist/arm64-${BINARY_NAME}

	# Build for linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LD_FLAGS} -o $(DIST_PATH)/linux64-$(BINARY_NAME) ./cmd/gnm/*.go
	@tar czvf ./dist/${BINARY_NAME}-linux64-${BUILD_VERSION}.tar.gz ./dist/linux64-${BINARY_NAME}
	
	# Build for windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${LD_FLAGS} -o $(DIST_PATH)/win64-$(BINARY_NAME) ./cmd/gnm/*.go
	@tar czvf ./dist/${BINARY_NAME}-win64-${BUILD_VERSION}.tar.gz ./dist/win64-${BINARY_NAME}

	@rm -rf ./$(DIST_PATH)/*-$(BINARY_NAME)