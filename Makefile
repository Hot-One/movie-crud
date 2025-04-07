CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

run: 
	go run ${APP_CMD_DIR}/main.go

swag-gen:
	echo ${REGISTRY}
	swag init -g api/router.go -o api/docs
