ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

up:
	docker-compose -f ${ROOT_DIR}/scrapper/docker-compose.yml up -d \
	&& cd ${ROOT_DIR}/api && ./vendor/bin/sail up -d

down:
	cd ${ROOT_DIR}/scrapper && docker-compose down \
	&& cd ${ROOT_DIR}/api && ./vendor/bin/sail down