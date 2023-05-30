package main

import (
	"os"
	"flag"
)

func main() {
	flag.Parse()
	command := flag.Arg(0)

	if command == "" {
		os.Exit(1)
	}

	switch command {
	case "up":
		migrateUp()
	default:
		os.Exit(0)
	}
}