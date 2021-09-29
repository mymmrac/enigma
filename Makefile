run: ## Run from source
	go run cmd/enigma/enigma.go

build: ## Build binary
	go build -o bin/enigma cmd/enigma/enigma.go

lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0

lint:
	$(shell go env GOPATH)/bin/golangci-lint run

.PHONY: run build lint-install lint
