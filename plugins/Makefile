#!/usr/bin/make -f

all: auth hello

auth: PACKAGE=./auth
auth: AOUT=auth.so

hello: PACKAGE=./hello
hello: AOUT=hello.so

TARGETS := auth hello

$(TARGETS):
	go build -buildmode=plugin -o $(AOUT) $(PACKAGE)
.PHONY: $(TARGETS)

clean:
	rm -f *.so

.PHONY: all clean build-plugin
