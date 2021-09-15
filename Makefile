ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

compose-build:
	export LOCAL_MACHINE_IP=${MACHINE_IP} \
	&& cd ${ROOT_DIR} && docker-compose build \
	&& cd ${ROOT_DIR}/api && ./vendor/bin/sail build

compose-up:
	export LOCAL_MACHINE_IP=${MACHINE_IP} \
	&& cd ${ROOT_DIR} && docker-compose up -d \
	&& cd ${ROOT_DIR}/api && ./vendor/bin/sail up -d

compose-down:
	cd ${ROOT_DIR} && docker-compose down \
	&& cd ${ROOT_DIR}/api && ./vendor/bin/sail down

laravel-sail:
	cd ${ROOT_DIR}/api && ./vendor/bin/sail $(call args,defaultstring)