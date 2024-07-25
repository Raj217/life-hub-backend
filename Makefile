build:
	go build -o bin/life_hub_backend cmd/life_hub_backend.go

run: build
	./bin/life_hub_backend

download:
	go mod download

test:
	go test ./...