.PHONY: migrate
migrate:
	go run tools/migrate/main.go

.PHONY: mock
mock:
	go generate ./...

.PHONY: run
run:
	go run cmd/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...