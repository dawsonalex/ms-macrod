GO_MODULE_PATH = github.com/dawsonalex/ms-macrod

# ROOT_DIR is the path of the makefile (including trailing slash)
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
PROJECT_PATH := $(ROOT_DIR:/=)
BIN_DIR = bin
BIN_NAME = macrod

# SHELL is set to a wrapper file that sources a helpers file and executes
# the command passed by Make.
SHELL := $(PWD)/shell

.PHONY: run # Run the server using go run
run:
	go run ${GO_MODULE_PATH}/cmd/...

.PHONY: build # build the server, default output to ${BIN_DIR}/${BIN_NAME}, embedding build info
build:
	go_build_with_version_info ${BIN_DIR}/${BIN_NAME} ${GO_MODULE_PATH} ${GO_MODULE_PATH}/cmd

.PHONY: rm-bin # remove the bin dir
rm-bin:
	rm -r ${PROJECT_PATH}/bin

.PHONY: help # Generate list of targets with descriptions
help:
	makeHelp
