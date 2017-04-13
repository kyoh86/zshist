package main

import (
	"log"
	"os"

	"github.com/kyoh86/zshist"
)

func main() {
	if err := zshist.Metafy(os.Stdin, os.Stdout); err != nil {
		log.Fatalf("failed to metafy (%s)", err.Error())
	}
}
