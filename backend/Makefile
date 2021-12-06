gen-automation-core-mock:
	@mockgen -source=internal/automation/interfaces/interfaces.go -destination=internal/automation/interfaces/mock/interfaces.go

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

gen-telegramclient-mock:
	@mockgen -source=internal/telegramClient/telegramClient.go -destination=internal/telegramClient/mock/telegramClient.go
