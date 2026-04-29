output := 'builds/'
projectName := 'CodeReport'

build:
	go build

run:
	go run

build-all: build-win-amd64 build-linux-amd64 build-win-arm build-linux-arm

build-win-amd64:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $(output)/$(projectName)-win-amd64.exe

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(output)/$(projectName)-linux-amd64

build-win-arm:
	GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -o $(output)/$(projectName)-win-arm.exe

build-linux-arm:
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o $(output)/$(projectName)-linux-arm

fmt: install-modules
	gofumpt -w .
	golangci-lint config verify
	golangci-lint run --fix


install-modules:
	go install mvdan.cc/gofumpt@v0.9.2
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4