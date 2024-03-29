GOBIN=$(shell pwd)/bin
GOFILES=cmd/main.go
GONAME=$(shell basename "$(PWD)")
PID=/tmp/go-$(GONAME).pid
# git tag --list '0.0.1' | xargs -I % echo "git tag -d %; git push --delete origin %" | sh

help: ## Display this help text
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


build: ## builds the binary
	@echo "Building $(GOFILES) to ./bin"
	go build -o bin/$(GONAME) cmd/main.go


migrate: ## installs all the dependencies
	go run cmd/main.go migrate


get: ## get all the dependencies
	@GOBIN=$(GOBIN) go get cmd/main.go

install: ## installs all the dependencies
	go install $(GOFILES)

run: ## runs the build
	go run $(GOFILES)

watch: ## This is for development. Restarts after every save
	@$(MAKE) restart &
	@fswatch -o . -e 'bin/.*' -ignore='Path vendor' | xargs -n1 -I{}  make restart

restart: clear stop clean build start

test: ## Run tests
	@go test -race  ./...


format: ## Format go code with goimports
	@go get golang.org/x/tools/cmd/goimports
	@goimports -l -w .

format-check: ## Check if the code is formatted
	@go get golang.org/x/tools/cmd/goimports
	@for i in $$(goimports -l .); do echo "[ERROR] Code is not formated run 'make format'" && exit 1; done

check: format-check ## Linting and static analysis
	@if grep -r --include='*.go' -E "fmt.Print|spew.Dump" *; then \
		echo "code contains fmt.Print* or spew.Dump function"; \
		exit 1; \
	fi

	@if test ! -e ./bin/golangci-lint; then \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh; \
	fi
	@./bin/golangci-lint run --timeout 180s -E gosec -E stylecheck -E revive -E goimports

start: ## start defending on localmachine
	@echo "Starting bin/$(GONAME)"
	@./bin/$(GONAME) & echo $$! > $(PID)

stop: ## stop defending on localmachine
	@echo "Stopping bin/$(GONAME) if it's running"
	@-kill `[[ -f $(PID) ]] && cat $(PID)` 2>/dev/null || true

clear: ## clear
	@clear

clean: ## clean
	@echo "Cleaning"
	@GOBIN=$(GOBIN) go clean

genproto: ## generate protobuf files
	@echo "Generating protobuf files"
	./script/generate-proto.sh

.PHONY: format check build get install run watch start stop restart clean