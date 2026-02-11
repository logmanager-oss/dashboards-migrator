.PHONY: build test lint vuln clean

BINARY_NAME=dashboards-migrator

build:
	CGO_ENABLED=0 go build -o $(BINARY_NAME) ./cmd/main.go

test:
	go test -v -race ./...

lint:
	golangci-lint run --timeout 5m

vuln:
	govulncheck -test ./...

clean:
	rm -f $(BINARY_NAME)
