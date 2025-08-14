
tmp/main: go.mod go.sum main.go
	mkdir -p tmp
	go build -o ./tmp/main

syft.json: tmp/main
	mkdir -p tmp
	syft ./tmp/main -o json=syft.ugly.json
	cat syft.ugly.json | jq . > tmp/syft.json
	rm syft.ugly.json

.PHONY: scan
scan: syft.json
	cat tmp/syft.json | jq  '.artifacts[].purl' -r | xargs -I{} grype {}

all: scan

# default goal
.DEFAULT_GOAL := all
