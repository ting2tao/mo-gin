#.PHONY: start build
#
#NOW = $(shell date -u '+%Y%m%d%I%M%S')
#
#RELEASE_VERSION = v1.0.0

APP_SERVER 		= server
APP_API 		= api
BIN_SERVER  	= ./build/${APP_SERVER}
BIN_API 		= ./build/${APP_API}
#GIT_COUNT 		= $(shell git rev-list --all --count)
#GIT_HASH        = $(shell git rev-parse --short HEAD)
#RELEASE_TAG     = $(RELEASE_VERSION).$(GIT_COUNT).$(GIT_HASH)

#all: run-server
#
#build: clean initConfig build-api build-server
#
build-api:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 packr build -ldflags "-w -s -X main.VERSION=$(RELEASE_TAG)" -o $(BIN_API) ./cmd/${APP_API}

#build-server:
#	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 packr build -ldflags "-w -s -X main.VERSION=$(RELEASE_TAG)" -o $(BIN_SERVER) ./cmd/${APP_SERVER}
#
#air-api:
#	air -c .air-api.toml
#
#air-server:
#	air -c .air-server.toml
#
run-server:
	@go run -ldflags "-X main.VERSION=$(RELEASE_TAG)" ./cmd/${APP_SERVER}/main.go

run-api:
	@go run -ldflags "-X main.VERSION=$(RELEASE_TAG)" ./cmd/${APP_API}/main.go

#migrate: wire-server
#	@go run ./cmd/migrate/main.go # 数据库迁移

wire-server:
	@wire gen ./app/server

wire-app:
	@wire gen ./internal/app

#clean:
#	@rm -rf $(BIN_SERVER)
#	@rm -rf $(BIN_API)
#	@rm -rf ./configs/config.ini
#
#initConfig: wire-server wire-api
#	@go run ./cmd/initConfig/main.go  # 初始化配置