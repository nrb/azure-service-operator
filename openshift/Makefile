# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL:=/usr/bin/env bash

.DEFAULT_GOAL:=help

GOPATH  := $(shell go env GOPATH)
GOARCH  := $(shell go env GOARCH)
GOOS    := $(shell go env GOOS)
GOPROXY := $(shell go env GOPROXY)
ifeq ($(GOPROXY),)
GOPROXY := https://proxy.golang.org
endif
export GOPROXY

# Active module mode, as we use go modules to manage dependencies
export GO111MODULE=on

# Kubebuilder.
export KUBEBUILDER_ENVTEST_KUBERNETES_VERSION ?= 1.29.0
export KUBEBUILDER_CONTROLPLANE_START_TIMEOUT ?= 60s
export KUBEBUILDER_CONTROLPLANE_STOP_TIMEOUT ?= 60s

# Directories
# Use the repository root, not the openshift directory.
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))/..))
BIN_DIR := $(ROOT_DIR)/bin
TOOLS_DIR := $(ROOT_DIR)/tools
TOOLS_BIN_DIR := $(abspath $(TOOLS_DIR)/bin)


# Envtest
SETUP_ENVTEST_VER := v0.0.0-20231012212722-e25aeebc7846
SETUP_ENVTEST_BIN := setup-envtest
SETUP_ENVTEST := $(abspath $(TOOLS_BIN_DIR)/$(SETUP_ENVTEST_BIN)-$(SETUP_ENVTEST_VER))
SETUP_ENVTEST_PKG := sigs.k8s.io/controller-runtime/tools/setup-envtest

KUBEBUILDER_ASSETS ?= $(shell $(SETUP_ENVTEST) use --use-env -p path $(KUBEBUILDER_ENVTEST_KUBERNETES_VERSION))

# TODO(nrb): Where is this set?
$(RELEASE_DIR):
	mkdir -p $(RELEASE_DIR)/

.PHONY: go-test
go-test: $(SETUP_ENVTEST) ## Run go tests.
	# TODO(nrb): This prorbably isn't the right command
	KUBEBUILDER_ASSETS="$(KUBEBUILDER_ASSETS)" go test ./... $(TEST_ARGS)

# TODO(nrb): Install envtest
 .PHONY: $(SETUP_ENVTEST_BIN)
$(SETUP_ENVTEST_BIN): $(SETUP_ENVTEST) ## Build a local copy of setup-envtest.

$(SETUP_ENVTEST): # Build setup-envtest from tools folder.
	#TODO(nrb): The GO_INSTALL script probably needs to be brought in, too.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) $(SETUP_ENVTEST_PKG) $(SETUP_ENVTEST_BIN) $(SETUP_ENVTEST_VER)

# TODO(nrb): Document e2e variables needed
# TODO(nrb): Write test wrapper



.PHONY: controller-generate-crds
controller-generate-crds:
	echo "TODO"

.PHONY: controller-run-kustomize-for-envtest
controller-run-kustomize-for-envtest:
	echo "TODO"


.PHONY: install-deps
install-deps:
	${ROOT_DIR}/.devcontainer/install-dependencies.sh

.PHONY: generate-crds
generate-crds:
	task controller:generate-crds


.PHONY: update-azure-rest-specs
update-azure-rest-specs:
	git submodule update --init
	cp v2/specs/azure-rest-api-specs/LICENSE openshift/azure-rest-api-specs/LICENSE
	cp -r v2/specs/azure-rest-api-specs/specification/** openshift/azure-rest-api-specs/specification/

# task: "controller:run-kustomize-for-envtest" started
# task: "controller:generate-kustomize" started
# task: "controller:generate-crds" started
# task: "controller:generate-types" started
# task: "controller:generate-genruntime-deepcopy" started
# task: "generator:build" started
# task: [generator:build] go build -ldflags "-X github.com/Azure/azure-service-operator/v2/internal/version.BuildVersion=v2.8.0-rev.46-dirty" -o ../../bin/aso-gen .
# task: Task "controller:generate-genruntime-deepcopy" is up to date
# task: "generator:build" finished
# task: Task "controller:generate-types" is up to date
# task: Task "controller:generate-crds" is up to date
# task: Task "controller:generate-kustomize" is up to date
# task: Task "controller:run-kustomize-for-envtest" is up to date

