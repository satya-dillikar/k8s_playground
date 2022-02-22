package main

import (
	"satya.com/burgerstore/dyn_client"
)

func main() {
	kubeConfigPath := "/Users/sdillikar/kind-basic.yaml"
	namespace := "default"
	dyn_client.DynamicMain(kubeConfigPath, namespace)

}
