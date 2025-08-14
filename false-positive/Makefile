all: build-and-scan

tmp/main: go.mod go.sum main.go
	go build -o ./tmp/main

syft.json: tmp/main
	syft ./tmp/main -o json=syft.ugly.json
	cat syft.ugly.json | jq . > syft.json
	rm syft.ugly.json

build-and-scan: syft.json
