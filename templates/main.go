package main

import (
	"fmt"
	"os"

	"github.com/sascha-andres/flag"
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

	displayHelp() // will exit with 0 if -help is passed

	return nil
}

func displayHelp() {
	if !help {
		return
	}
	flag.PrintDefaults()
	os.Exit(0)
}
