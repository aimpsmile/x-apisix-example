BUILD_WEB_NAME:=aimgo.web.v1.test
BUILD_WEBSOCKET_NAME:=aimgo.web.v1.websocket
BUILD_HTTP2_NAME:=go.http2.v1.passport
BUILD_XAPI_NAME:=x-apisix
MKFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
EXEC_ENV:=AIM_CONFIG=${MKFILE_DIR}apisix/config AIM_ENV=local MICRO_REGISTRY=etcd MICRO_REGISTRY_ADDRESS=etcd.service:2379

.PHONY: test-install test-exec websocket-install websocket-exec xapisix-monitor
web-install:
	go build -i -o ${GOPATH}/bin/${BUILD_WEB_NAME} -ldflags "-w -s"  ./web/main.go

web-exec:
	${EXEC_ENV} ${BUILD_WEB_NAME}

websocket-install:
	go build -i -o ${GOPATH}/bin/${BUILD_WEBSOCKET_NAME} -ldflags "-w -s"  ./websocket/main.go

websocket-exec:
	${EXEC_ENV} ${BUILD_WEBSOCKET_NAME}

xapisix-install:
	go build -i -o ${GOPATH}/bin/${BUILD_XAPI_NAME} -ldflags "-w -s"  ./apisix/main.go ./apisix/plugins.go

xapisix-exec:
	  MICRO_SERVER_VERSION=v0.6.1 AIM_GATEWAY=apisix MICRO_SERVER_NAME=${BUILD_XAPI_NAME} ${EXEC_ENV} ${BUILD_XAPI_NAME} monitor


grpc-install:
	go build -i -o ${GOPATH}/bin/${BUILD_HTTP2_NAME} -ldflags "-w -s"  ./grpc/main.go

grpc-exec:
	${EXEC_ENV} ${BUILD_HTTP2_NAME}
