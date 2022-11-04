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
	GOOS=windows GOARCH=arm go build -ldflags "-s -w" -o $(output)/$(projectName)-win-arm.exe

build-linux-arm:
	GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o $(output)/$(projectName)-linux-arm