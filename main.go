package main

import (
	"multi-kubectl/pkg/kubectl"
	"os"
)

func main() {
	kubectl.RunCommand(os.Args)
}
