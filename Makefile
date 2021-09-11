ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

build:
	docker-compose -f ${ROOT_DIR}/docker-compose.yml build \
	&& cd ${ROOT_DIR}/api && ./vendor/bin/sail build

up:
	docker-compose -f ${ROOT_DIR}/docker-compose.yml up -d \
	&& cd ${ROOT_DIR}/api && ./vendor/bin/sail up -d

down:
	cd ${ROOT_DIR} && docker-compose down \
	&& cd ${ROOT_DIR}/api && ./vendor/bin/sail down