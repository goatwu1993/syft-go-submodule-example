
tmp/main: go.mod go.sum main.go
	go build -o ./tmp/main

syft.json: tmp/main
	syft ./tmp/main -o json=syft.ugly.json
	cat syft.ugly.json | jq . > syft.json
	rm syft.ugly.json

.PHONY: scan
scan: syft.json
	cat syft.json | jq  '.artifacts[].purl' -r | xargs -I{} grype {}
