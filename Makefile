module_name := $(shell grep '^module ' go.mod | awk '{print $$2}')

test:
	go clean -testcache
	go test -v ./...
.PHONY: test

build:
	@mkdir -p bin
	@echo "Building ${module_name}..."
	go build -o bin/$(module_name) ./cmd/main.go || exit 1
	.

.PHONY: build

clean:
	go clean
	rm -rf bin/
.PHONY: clean

