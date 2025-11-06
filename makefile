build:
	go build -o server ./cmd/api/main.go
clean:
	go mod tidy
