package main

import (
	"fmt"
	"os"

	"github.com/go-viper/mapstructure/v2"
)

func main() {
	type config struct {
		Foo string `mapstructure:"foo"`
	}
	var result config
	type otherConfig struct {
		Bar string `mapstructure:"bar"`
		Foo string `mapstructure:"foo"`
	}
	original := otherConfig{
		Bar: "baz",
		Foo: "qux",
	}
	err := mapstructure.WeakDecode(original, &result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Decoded config: %+v\n", result)
}
