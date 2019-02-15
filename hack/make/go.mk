# ----------------------------------------------------------------------------
# global

SHELL = /usr/bin/env bash
GO_PATH = $(shell go env GOPATH)

PKG := $(subst $(GO_PATH)/src/,,$(CURDIR))
GO_PKGS := $(shell go list ./... | grep -v -e '.pb.go')
GO_APP_PKGS := $(shell go list -f '{{if and (or .GoFiles .CgoFiles) (ne .Name "main")}}{{.ImportPath}}{{end}}' ${PKG}/...)
GO_TEST_PKGS := $(shell go list -f='{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}' ./...)
GO_VENDOR_PKGS := $(shell go list -f '{{if and (or .GoFiles .CgoFiles) (ne .Name "main")}}./vendor/{{.ImportPath}}{{end}}' ./vendor/...)

GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_UNTRACKED_CHANGES:= $(shell git status --porcelain --untracked-files=no)
ifneq ($(GIT_UNTRACKED_CHANGES),)
	GIT_COMMIT := $(GIT_COMMIT)-dirty
endif

CGO_ENABLED ?= 0
GO_LDFLAGS:=-ldflags "-w"
GO_LDFLAGS_STATIC:=-ldflags "-w -extldflags -static"
GO_BUILD_TAGS := osusergo
GO_FLAGS ?= -tags '$(GO_BUILD_TAGS)'

ifneq ($(wildcard go.mod),)  # exist go.mod file
ifeq ($(CI),)  # $CI is empty
	GO_FLAGS+=-mod=vendor
else
	GO_FLAGS+=-mod=readonly
endif
endif

GO_TEST ?= go test
GO_TEST_FUNC ?= .
GO_TEST_FLAGS ?=
GO_BENCH_FUNC ?= .
GO_BENCH_FLAGS ?= -benchmem

IMAGE_REGISTRY := gcr.io/container-image

# ----------------------------------------------------------------------------
# defines

GOPHER = ""
define target
@printf "$(GOPHER)  \\033[32m$(patsubst ,$@,$(1))\\033[0m\\n"
endef

# ----------------------------------------------------------------------------
# targets

.PHONY: pkg/install
pkg/install:
	CGO_ENABLED=$(CGO_ENABLED) go install -v ${GO_APP_PKGS}

## test, bench and coverage

.PHONY: test
test:  ## Run the package test with checks race condition.
	$(call target)
	$(GO_TEST) -v -race $(strip $(GOFLAGS)) -run=$(GO_TEST_FUNC) $(GO_TEST_PKGS)

.PHONY: bench
bench:  ## Take a package benchmark.
	$(call target)
	$(GO_TEST) -v $(strip $(GOFLAGS)) -run='^$$' -bench=$(GO_BENCH_FUNC) -benchmem $(GO_TEST_PKGS)

.PHONY: bench/race
bench/race:  ## Take a package benchmark with checks race condition.
	$(call target)
	$(GO_TEST) -v -race $(strip $(GOFLAGS)) -run='^$$' -bench=$(GO_BENCH_FUNC) -benchmem $(GO_TEST_PKGS)

.PHONY: coverage
coverage:  ## Take test coverage.
	$(call target)
	$(GO_TEST) -v -race $(strip $(GOFLAGS)) -covermode=atomic -coverpkg=$(PKG)/... -coverprofile=coverage.out $(GO_PKGS)

$(GO_PATH)/bin/go-junit-report:
	@GO111MODULE=off go get -u github.com/jstemmer/go-junit-report

.PHONY: cmd/go-junit-report
cmd/go-junit-report: $(GO_PATH)/bin/go-junit-report  # go get 'go-junit-report' binary

.PHONY: coverage/junit
coverage/junit: cmd/go-junit-report  ## Take test coverage and output test results with junit syntax.
	$(call target)
	@mkdir -p test-results
	$(GO_TEST) -v -race $(strip $(GO_FLAGS)) -covermode=atomic -coverpkg=$(PKG)/... -coverprofile=coverage.out $(GO_PKGS) 2>&1 | tee /dev/stderr | go-junit-report -set-exit-code > test-results/report.xml


## lint

.PHONY: lint
lint: lint/golangci-lint  ## Run all linters.

$(GO_PATH)/bin/golangci-lint:
	@GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: cmd/golangci-lint
cmd/golangci-lint: $(GO_PATH)/bin/golangci-lint  # go get 'golangci-lint' binary

.PHONY: golangci-lint
lint/golangci-lint: cmd/golangci-lint .golangci.yml  ## Run golangci-lint.
	$(call target)
	@golangci-lint run ./...


## mod

.PHONY: mod/init
mod/init:
	$(call target)
	@GO111MODULE=on go mod init || true

.PHONY: mod/goget
mod/goget:  ## Update module and go.mod.
	$(call target)
	@GO111MODULE=on go get -u -m -v -x

.PHONY: mod/tidy
mod/tidy:
	$(call target)
	@GO111MODULE=on go mod tidy -v

.PHONY: mod/vendor
mod/vendor:
	$(call target)
	@GO111MODULE=on go mod vendor -v

.PHONY: mod/graph
mod/graph:
	$(call target)
	@GO111MODULE=on go mod graph

.PHONY: mod/clean
mod/clean:
	$(call target)
	@$(RM) go.mod go.sum
	@$(RM) -r vendor

.PHONY: mod/install
mod/install: mod/tidy mod/vendor
	$(call target)
	@GO111MODULE=off go install -v $(GO_VENDOR_PKGS) || GO111MODULE=on go install -mod=vendor -v $(GO_VENDOR_PKGS)

.PHONY: mod
mod: mod/init mod/tidy mod/vendor mod/install  ## Updates the vendoring directory via go mod.
	@sed -i ':a;N;$$!ba;s|go 1\.12\n\n||g' go.mod

.PHONY: mod/update
mod/update: mod/goget mod/tidy mod/vendor mod/install  ## Updates all vendor packages.


## miscellaneous

.PHONY: boilerplate/go/%
boilerplate/go/%: BOILERPLATE_PKG_DIR=$(shell printf $@ | cut -d'/' -f3- | rev | cut -d'/' -f2- | rev)
boilerplate/go/%: BOILERPLATE_PKG_NAME=$(if $(findstring main,$@),main,$(shell printf $@ | rev | cut -d/ -f2 | rev))
boilerplate/go/%: hack/boilerplate/boilerplate.go.txt  ## Create go file from boilerplate.go.txt
	@if [ ${BOILERPLATE_PKG_DIR} != "*.go" ] && [ ! -d ${BOILERPLATE_PKG_DIR} ]; then mkdir -p ${BOILERPLATE_PKG_DIR}; fi
	@cat hack/boilerplate/boilerplate.go.txt <(printf "package ${BOILERPLATE_PKG_NAME}\\n") > $*
	@sed -i "s|YEAR|$(shell date '+%Y')|g" $*


.PHONY: AUTHORS
AUTHORS:  ## Creates AUTHORS file.
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@printf "$(shell git log --format="\n%aN <%aE>" | LC_ALL=C.UTF-8 sort -uf)" >> $@


.PHONY: clean
clean:  ## Cleanup any build binaries or packages.
	$(call target)
	@$(RM) *.out *.test *.prof trace.log


.PHONY: help
help:  ## Show make target help.
	@perl -nle 'BEGIN {printf "Usage:\n  make \033[33m<target>\033[0m\n\nTargets:\n"} printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 if /^([a-zA-Z\/_-].+)+:.*?\s+## (.*)/' ${MAKEFILE_LIST}
