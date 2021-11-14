SHELL=/bin/bash
export GOBIN := $(PWD)/bin
GO_PATH:=$(shell go env GOPATH)
export PATH:=$(GOBIN):$(GO_PATH):$(PATH)

PROJ_DIR:=$(shell pwd)

# Docker settings
ENABLE_DOCKER_BUILD_KIT:=DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1
DOCKER_BUILD_BASE:=  $(ENABLE_DOCKER_BUILD_KIT) docker-compose -f docker-compose.yml

## makeのヘルプ表示
help: setup-help
	@$(GOBIN)/make2help $(MAKEFILE_LIST)

setup-help:
	@if [[  ! -x $(GOBIN)/make2help ]]; then \
		$(GOMOD) go install github.com/Songmu/make2help/cmd/make2help ; \
	fi

## 最初の準備するために全てのpackagesをinstallする
bootstrap: install-go install-tools update-module

## .go-versonに記載されているバージョンのgoをインストール
install-go:
	@./tools/install_go.sh

## toolをinstallする
install-tools:
	@$(GOMOD) /bin/bash ./tools/install_tools.sh

## 依存するpackageを最新版に更新する。
update-module:
	$(GOMOD) go mod tidy

## lintをチェックする
lint:
	@$(GOMOD) $(GOBIN)/golangci-lint run
	@$(GOMOD) go vet -vettool=$(GOBIN)/todocomment -issue "https://github.com/nhannt315/DontMiss2Man/issues" ./...

## go testを実行
test:
	go test -race ./...

## go testを実行・カバレッジを出力
test-with-coverage:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

##
build-migration:
	$(DOCKER_BUILD_BASE) build migration

migration-create:
	sh ./deployments/migration/migration.sh create $(name)

## [docker/local] マイグレーション(up) ビルド/スタート
migration-up: start-db
	$(DOCKER_BUILD_BASE) run --rm migration /db/migration.sh up

## [docker/local] マイグレーション(down) ビルド/スタート
migration-down: start-db
	$(DOCKER_BUILD_BASE) run --rm migration /db/migration.sh down

## [docker/local] マイグレーション(status) ビルド/スタート
migration-status: start-db
	$(DOCKER_BUILD_BASE) run --rm migration /db/migration.sh status

## DBスタート
start-db:
	$(DOCKER_BUILD_BASE) up -d db
	./deployments/wait_service_up.sh --service db --text "port: 3306  MySQL"

## DBストップ
stop-db:
	$(DOCKER_BUILD_BASE) stop db

clear-db-data: stop-db
	rm -fr deployments/db/data/*

reset-db: clear-db-data
	$(DOCKER_BUILD_BASE) rm -s -v -f db
	make start-db
	make migration-up
