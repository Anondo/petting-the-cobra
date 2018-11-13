package main

import (
	"cliapp/cmd"
	"log"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
