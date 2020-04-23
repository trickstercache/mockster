#
# Copyright 2018 Comcast Cable Communications Management, LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

DEFAULT: build

GO             ?= go
GOFMT          ?= $(GO)fmt
FIRST_GOPATH   := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
MOCKSTER_MAIN := cmd/mockster
MOCKSTER      := $(FIRST_GOPATH)/bin/mockster
PROGVER        := $(shell grep 'applicationVersion = ' $(MOCKSTER_MAIN)/main.go | awk '{print $$3}' | sed -e 's/\"//g')
BUILD_TIME     := $(shell date -u +%FT%T%z)
GIT_LATEST_COMMIT_ID     := $(shell git rev-parse HEAD)
GO_VER         := $(shell go version | awk '{print $$3}')
LDFLAGS=-ldflags "-s -X main.applicationBuildTime=$(BUILD_TIME) -X main.applicationGitCommitID=$(GIT_LATEST_COMMIT_ID) -X main.applicationGoVersion=$(GO_VER) -X main.applicationGoArch=$(GOARCH)"
IMAGE_TAG      ?= latest
IMAGE_ARCH     ?= amd64
GOARCH         ?= amd64
TAGVER         ?= unspecified

.PHONY: validate-app-version
validate-app-version:
	@if [ "refs/tags/v$(PROGVER)" != $(TAGVER) ]; then\
		(echo "mismatch between TAGVER '$(TAGVER)' and applicationVersion '$(PROGVER)'"; exit 1);\
	fi

.PHONY: go-mod-vendor
go-mod-vendor:
	$(GO) mod vendor

.PHONY: go-mod-tidy
go-mod-tidy:
	$(GO) mod tidy

.PHONY: build
build: go-mod-tidy go-mod-vendor
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GO) build $(LDFLAGS) -o ./OPATH/mockster -a -v $(MOCKSTER_MAIN)/*.go

.PHONY: install
install:
	$(GO) install -o $(MOCKSTER) $(PROGVER)

.PHONY: release
release: validate-app-version clean go-mod-tidy go-mod-vendor release-artifacts # docker docker-release

.PHONY: release-artifacts
release-artifacts: validate-app-version
	GOOS=darwin   GOARCH=amd64 $(GO) build $(LDFLAGS) -o ./OPATH/mockster-$(PROGVER).darwin-amd64       -a -v $(MOCKSTER_MAIN)/*.go && tar cvfz ./OPATH/mockster-$(PROGVER).darwin-amd64.tar.gz  ./OPATH/mockster-$(PROGVER).darwin-amd64
	GOOS=linux    GOARCH=amd64 $(GO) build $(LDFLAGS) -o ./OPATH/mockster-$(PROGVER).linux-amd64        -a -v $(MOCKSTER_MAIN)/*.go && tar cvfz ./OPATH/mockster-$(PROGVER).linux-amd64.tar.gz   ./OPATH/mockster-$(PROGVER).linux-amd64
	GOOS=linux    GOARCH=arm64 $(GO) build $(LDFLAGS) -o ./OPATH/mockster-$(PROGVER).linux-arm64        -a -v $(MOCKSTER_MAIN)/*.go && tar cvfz ./OPATH/mockster-$(PROGVER).linux-arm64.tar.gz   ./OPATH/mockster-$(PROGVER).linux-arm64
	GOOS=windows  GOARCH=amd64 $(GO) build $(LDFLAGS) -o ./OPATH/mockster-$(PROGVER).windows-amd64.exe  -a -v $(MOCKSTER_MAIN)/*.go && tar cvfz ./OPATH/mockster-$(PROGVER).windows-amd64.tar.gz ./OPATH/mockster-$(PROGVER).windows-amd64.exe

.PHONY: style
style:
	! gofmt -d $$(find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

.PHONY: test
test:
	@go test -v -coverprofile=.coverprofile ./...

.PHONY: test-cover
test-cover: test
	$(GO) tool cover -html=.coverprofile

.PHONY: clean
clean:
	rm -rf ./mockster ./OPATH
