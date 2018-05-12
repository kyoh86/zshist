package main

import (
	"log"
	"os"

	"github.com/kyoh86/zshist"
)

func main() {
	if err := zshist.Unmetafy(os.Stdin, os.Stdout); err != nil {
		log.Fatalf("failed to unmetafy (%s)", err.Error())
	}
}
