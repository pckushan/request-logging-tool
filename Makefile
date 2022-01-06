build:
	@go build -o myhttp

tests:
	@go test ./... -count=1
