override APP_NAME=api
override GO_VERSION=1.22

GOOS?=$(shell go env GOOS || echo linux)
GOARCH?=$(shell go env GOARCH || echo amd64)
CGO_ENABLED?=1

.PHONY: all
all: cleanup vendor

.PHONY: cleanup
cleanup:
	@rm ${PWD}/bin/${APP_NAME} || true
	@rm -r ${PWD}/vendor || true

.PHONY: vendor
vendor:
	@go mod vendor

.PHONY: build
build:
	set CGO_ENABLED=1 && set GOOS=linux && set GOARCH=amd64 && go build -ldflags=-linkmode=external \
		-mod vendor \
		-o ./bin/${APP_NAME}.a \
		-buildmode=c-archive ./pkg/api
