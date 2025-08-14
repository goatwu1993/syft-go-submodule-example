# syft-go-submodule-example

## How to run scan

```bash
syft . -o json | jq  '.artifacts[].purl' -r | grep 'pkg:golang/github.com/hashicorp/vault'
```