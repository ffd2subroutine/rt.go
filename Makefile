test:
	go test ./...
build:
	go build -ldflags="-s -w" -o bin/rt cmd/rt/main.go
