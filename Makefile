build:
	@go build main.go

run: build
	@./main -filename="instructions.csv"

test:
	@go get github.com/stretchr/testify/assert
	@go test ./...
