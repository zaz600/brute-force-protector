build: build-cli build-server

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.30.0

lint:
	golangci-lint run ./...

test:
	go fmt ./internal/...
	go vet ./internal/...
	go test -v ./internal/...
	go test -v -race ./internal/...
	go test -gcflags=-l -count=1 -timeout=30s -bench=. -run=^$  ./internal/...

	go test -cover ./internal/... | grep coverage

build-server:
	go build -o ./bin/bp-server ./cmd/bp-server

build-cli:
	go build -o ./bin/bp-cli ./cmd/bp-cli

release: build test lint

run:
	docker-compose -f build/docker-compose.yml -p bruteforce-protector up -d --build

run-log:
	docker-compose -f build/docker-compose.yml -p bruteforce-protector up --build

stop:
	docker-compose -f build/docker-compose.yml -p bruteforce-protector down

generate:
	protoc --go_out=internal/grpc/ --go-grpc_out=internal/grpc/ api/api.proto

itest:
	docker-compose -f build/docker-compose-itest.yml -p bruteforce-protector-itest up --build -d
	docker-compose -f build/docker-compose-itest.yml -p bruteforce-protector-itest logs --follow itests
	docker-compose -f build/docker-compose-itest.yml -p bruteforce-protector-itest down

itest-stop:
	docker-compose -f build/docker-compose-itest.yml -p bruteforce-protector-itest down

itest-local:
	go test -v ./tests

