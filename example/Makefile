# based in part on: https://povilasv.me/exposing-go-modules-to-prometheus/
PROGRAM := example
PACKAGE := github.com/cisco-cx/informer/$(PROGRAM)
GO_VERSION :=$(shell go version)
LICENSE := Apache-2.0
URL     := https://$(PACKAGE)
DATE := $(shell date +%FT%T%z)
USER := $(shell whoami)
GIT_HASH := $(shell git --no-pager describe --tags --always)
BRANCH := $(shell git branch | grep '*' | cut -d ' ' -f2)

LDFLAGS := -s
LDFLAGS += -X "$(PACKAGE).Program=$(PROGRAM)"
LDFLAGS += -X "$(PACKAGE).License=$(LICENSE)"
LDFLAGS += -X "$(PACKAGE).URL=$(URL)"
LDFLAGS += -X "$(PACKAGE).BuildUser=$(USER)"
LDFLAGS += -X "$(PACKAGE).BuildDate=$(DATE)"
LDFLAGS += -X "$(PACKAGE).Version=$(GIT_HASH)"
LDFLAGS += -X "$(PACKAGE).Revision=$(GIT_HASH)"
LDFLAGS += -X "$(PACKAGE).Branch=$(BRANCH)"

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '$(LDFLAGS)' ./cmd/$(PROGRAM)
