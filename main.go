package main

import (
	"os"

	"github.com/alecthomas/kingpin"
)

// nolint
var (
	version = "snapshot"
	commit  = "snapshot"
	date    = "snapshot"
)

func main() {
	app := kingpin.New("zshist", "A tool for Encoding/Decoding .zsh_history file.").Version(version).Author("kyoh86")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	}
}
