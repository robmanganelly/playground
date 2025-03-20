.DEFAULT_GOAL :=build
.PHONY: fmt lint vet tidy vendor build

tidy:
	go mod tidy
vendor: tidy
	go mod vendor
fmt: vendor
	go fmt ./...
lint: fmt
	golint ./...
vet: fmt
	go vet ./...
build: vet
	mkdir -p build
	go build -o build/hello hello.go
