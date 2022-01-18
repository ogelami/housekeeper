BINARY := housekeeper
CONFIG_FILE ?= housekeeper.conf

SYSCONFDIR ?= $(PREFIX)/etc
SBINDIR := $(PREFIX)/usr/sbin

export GOBIN=$(PWD)/bin

.PHONY: all build dep install clean

all: dep build

build: main.go
	go build -mod vendor -ldflags "-s -X main.SYSCONF_PATH=${SYSCONFDIR}" -o $(GOBIN)/$(BINARY)

dep:
	go mod vendor

run: main.go
	HOUSEKEEPER_CONFIGURATION_PATH=housekeeper.conf go run -mod vendor main.go

install:
	mkdir -p $(SYSCONFDIR) $(SBINDIR)
	cp -i $(CONFIG_FILE).template $(CONFIG_PATH)
	cp $(GOBIN)/$(BINARY) $(SBINDIR)/$(BINARY)

clean:
	rm -rf bin/*
	rm -rf vendor
