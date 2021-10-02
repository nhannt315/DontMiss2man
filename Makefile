SHELL=/bin/bash
PROJ_DIR:=$(shell pwd)

# Docker settings
ENABLE_DOCKER_BUILD_KIT:=DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1
DOCKER_BUILD_BASE:=  $(ENABLE_DOCKER_BUILD_KIT) docker-compose -f docker-compose.yml

start-db:
	$(DOCKER_BUILD_BASE) up -d db

read-env: 
	bash -c "export $(sed 's/[[:blank:]]//g; /^#/d' ./api/.env.local | xargs)"
