BINARY := housekeeper
CONFIG_FILE ?= housekeeper.conf

GOPATH := $(PWD)
GOBIN := $(GOPATH)/bin
SYSCONFDIR ?= $(PREFIX)/etc
CONFIG_PATH := $(SYSCONFDIR)/$(CONFIG_FILE)
SBINDIR := $(PREFIX)/usr/sbin

.PHONY: all build dep install clean

all : dep build

build : main.go
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -ldflags "-s -X main.CONFIGURATION_PATH=${CONFIG_PATH}" -o $(GOBIN)/$(BINARY)

dep:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -d

install:
	mkdir -p $(SYSCONFDIR) $(SBINDIR) $(LIBDIR)
	cp -i $(CONFIG_FILE).template $(CONFIG_PATH)
	cp $(GOBIN)/$(BINARY) $(SBINDIR)/$(BINARY)
	cp $(GOBIN)/*.so $(LIBDIR)

clean:
	go clean
	rm -rf $(GOBIN)/*

