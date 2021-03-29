TAG=$(shell cat .release | cut -d'=' -f2)
DEVELOPMENT_TAG=$(shell cat .development)
BUILD_PATH="$(shell go env GOPATH)/bin/wabacli"
.DEFAULT_GOAL := build

build: system-check
	@echo "starting build at $(BUILD_PATH) for tag $(DEVELOPMENT_TAG)"
	@cd cmd/ && GOOS_VAL=$(shell go env GOOS) GOARCH_VAL=$(shell go env GOARCH) go build -ldflags="-X main.BuildVersion=$(DEVELOPMENT_TAG)" -o $(BUILD_PATH) main.go
	@echo "build successful"

install: system-check
	@if [ ! -f $(BUILD_PATH) ] ; then echo "binaries does not exits at $(BUILD_PATH)"; exit 1; fi;
	@if [[ "$(shell go env GOOS)" == "darwin" ]]; then echo "copying binaries to install path" && sudo cp "${BUILD_PATH}" /usr/local/bin/; fi;

release: system-check test release-pre-check tag
	@echo "creating release $(TAG)"
	@goreleaser release

release-pre-check:
	@echo "running pre checks"
	@if [ -n "$(shell git tag | grep $(TAG))" ] ; then echo "ERROR: Tag '$(TAG)' already exits" && exit 1; fi;
	@if [ -z "$(shell git remote -v)" ] ; then echo "ERROR: no remote to push tag" && exit 1; fi;
	@if [ -z "$(shell git config user.email)" ] ; then echo 'ERROR: Unable to detect git credentials' && exit 1 ; fi
	@if [ -d dist ] ; then echo "deleting previously generated dist files" && rm -rf dist; fi

tag: update-readme
	@echo "creating tag $(TAG)"
	@git add .release README.md .goreleaser.yml cmd/ pkg/ config/ Makefile
	@git commit -m "Release $(TAG)"
	@git tag $(TAG)
	@git push origin $(TAG)

update-readme:
	@echo "updated README.md release tag to $(TAG)"
	@sed -i "" "s~release-v.*-blue)~release-${TAG}-blue)~" README.md

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