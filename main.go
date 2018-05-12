package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: zshist <command>")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "commands:")
	fmt.Fprintln(os.Stderr, "\tzshist unmetafy\t: Unmetafy(decode) for zsh_history")
	fmt.Fprintln(os.Stderr, "\tzshist metafy\t: Metafy(encode) for zsh_history")
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
		os.Exit(1)
	}
	switch flag.Arg(0) {
	case "unmetafy", "metafy":
		cmd := exec.Command("zshist-"+flag.Arg(0), flag.Args()[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	default:
		usage()
		os.Exit(1)
	}
}
