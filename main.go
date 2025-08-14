package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api/auth/kubernetes"
)

func main() {
	_, err := kubernetes.NewKubernetesAuth("customrole")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating KubernetesAuth: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("KubernetesAuth created successfully")
}