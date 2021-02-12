export

GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=cmd/$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")
PID=/tmp/go-$(GONAME).pid

## Display this help text
help:
	$(info Available Targets)
	@awk '/^[a-zA-Z\-\_0-9]+:/ {                    \
		nb = sub( /^## /, "", helpMsg );              \
		if(nb == 0) {                                 \
		helpMsg = $$0;                              \
		nb = sub( /^[^:]*:.* ## /, "", helpMsg );   \
		}                                             \
		if (nb)                                       \
		printf "\033[1;31m%-" width "s\033[0m %s\n", $$1, helpMsg;   \
	}                                               \
	{ helpMsg = $$0 }'                              \
	$(MAKEFILE_LIST) | column -ts:

## build the repo
build:
	@echo "Building $(GOFILES) to ./bin"
	go build -o bin/$(GONAME) cmd/main.go

## get all the dependencies
get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .
## installs all the dependencies
install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)
## runs the build
run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)
## This is for development. Restarts after every save
watch:
	@$(MAKE) restart &
	@fswatch -o . -e 'bin/.*' -ignore='Path vendor' | xargs -n1 -I{}  make restart

restart: clear stop clean build start

start:
	@echo "Starting bin/$(GONAME)"
	@./bin/$(GONAME) & echo $$! > $(PID)

stop:
	@echo "Stopping bin/$(GONAME) if it's running"
	@-kill `[[ -f $(PID) ]] && cat $(PID)` 2>/dev/null || true

clear:
	@clear

clean:
	@echo "Cleaning"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

.PHONY: build get install run watch start stop restart clean