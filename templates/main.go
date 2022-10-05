package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help bool
)

func init() {
	flag.BoolVar(&help, "help", false, "print help")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	help() // will exit with 0 if -help is passed

	return nil
}

func help() {
	if !help {
		return
	}
	flag.PrintDefaults()
	os.Exit(0)
}
