package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/sdk/helper/errutil"
)

func main() {
	var err error
	err = &errutil.InternalError{
		Err: "This is an internal error",
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating KubernetesAuth: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("KubernetesAuth created successfully")
}
