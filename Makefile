OUTPUT = bin/
GOCMD = go

.DEFAULT_GOAL := build

.PHONY: test

create_out_dir:
	mkdir -p $(OUTPUT)
test:
	$(GOCMD) test ./...

.PHONY: clean
clean:
	rm -rf $(OUTPUT)

.PHONY: install
install:
	$(GOCMD) get ./...

build: create_out_dir
	$(GOCMD) build -o $(OUTPUT) ./...

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda: create_out_dir
	GOOS=linux GOARCH=amd64 $(MAKE) all

.PHONY: build
all: clean lambda

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor

fmt:
	$(GOCMD) fmt ./...