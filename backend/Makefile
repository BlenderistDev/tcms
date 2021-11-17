gen-automation-core-mock:
	@mockgen -source=automation/core/core.go -destination=automation/core/mock.go -package=core

test:
	@go test ./... -race -count=1 -cover -coverprofile=coverage.txt && go tool cover -func=coverage.txt | tail -n1 | awk '{print "Total test coverage: " $$3}'
	@rm coverage.txt

format:
	@go fmt ./...
