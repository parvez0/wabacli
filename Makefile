TAG=$(shell cat .release | cut -d'=' -f2)
BUILD_PATH="$(shell go env GOPATH)/bin/wabactl"
.DEFAULT_GOAL := build

build: system-check
	@echo "starting build at $(BUILD_PATH)"
	@cd cmd/ && GOOS_VAL=$(shell go env GOOS) GOARCH_VAL=$(shell go env GOARCH) go build -o $(BUILD_PATH) main.go
	@echo "build successful"

install: system-check
	@if [ ! -f $(BUILD_PATH) ] ; then echo "binaries does not exits at $(BUILD_PATH)"; exit 1; fi;
	@if [ "$(go env GOOS)" == "darwin" ] ; then cp $(BUILD_PATH) /Users/$($whoami)/bin/wabactl; fi;

release: system-check test
	@echo "creating a release"

system-check:
	@echo "initializing system check"
	@if [ -z "$(shell go env GOOS)" ] || [ -z "$(shell go env GOARCH)" ] ;\
	 then \
   		echo "system info couldn't be determined" && exit 1 ; \
   	 else \
   	    echo "Go System: $(shell go env GOOS)" ; \
   	    echo "GO Arch: $(shell go env GOARCH)" ; \
   	    echo "system check passed" ;\
   	 fi ;

test:
	@echo "starting unit test"
	@go test ./pkg/tests