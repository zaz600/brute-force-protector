build: build-cli build-server

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.30.0

lint:
	golangci-lint run ./...

test:
	go test -v -race ./internal/...
	go test -gcflags=-l -count=1 -timeout=30s -bench=. -run=^$  ./internal/...

build-server:
	go build -o ./bin/bp-server ./cmd/bp-server

build-cli:
	go build -o ./bin/bp-cli ./cmd/bp-cli

release: build test lint

run:
	docker-compose -f docker-compose.yml -p bruteforce-protector up -d

stop:
	docker-compose -f docker-compose.yml -p bruteforce-protector down