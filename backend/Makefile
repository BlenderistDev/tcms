gen-automation-core-mock:
	@mockgen -source=internal/automation/core/core.go -destination=internal/automation/core/core_mock.go -package=core

test:
	@go test ./... -race -count=1 -cover -coverprofile=coverage.txt && go tool cover -func=coverage.txt | tail -n1 | awk '{print "Total test coverage: " $$3}'
	@rm coverage.txt

format:
	@go fmt ./...

run:
	@go run cmd/tcms/main.go

build:
	@mkdir -p ./bin
	@go build -o bin/tcms ./cmd/tcms
