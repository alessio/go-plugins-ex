#!/usr/bin/make -f

PLUGINS_DIR ?= $(CURDIR)/build/

go-plugins-ex:
	go build

all: go-plugins-ex modules

modules:
	$(MAKE) -C plugins/

install: modules
	go install
	install plugins/*.so $(shell go env GOPATH)/bin

clean:
	rm -rf $(CURDIR)/build/
	$(MAKE) -C plugins/ clean

.PHONY: all modules install clean
