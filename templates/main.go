package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sascha-andres/flag"
)

const envPrefix string = "TOOL"

var (
	help bool
)

func init() {
	log.SetPrefix(fmt.Sprintf("[%s] ", envPrefix))
	log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)

	flag.SetEnvPrefix(envPrefix)
	flag.BoolVarWithoutEnv(&help, "help", false, "print help")
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("error running %s: %s", envPrefix, err)
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
