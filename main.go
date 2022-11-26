package main

import (
	"k8s.io/client-go/rest"
)

func main() {
	_, err := rest.InClusterConfig()
	if err != nil {
		return
	}
}
