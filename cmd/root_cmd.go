package cmd

import (
	"log"
	"petting-the-cobra/conf"

	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	rootCmnd := cobra.Command{
		Use: "example",
		Run: run,
	}

	rootCmnd.Flags().IntP("port", "p", 0, "the port to do things on")
	rootCmnd.Flags().StringP("config", "c", "", "the configuration file")

	return &rootCmnd
}

func run(cmd *cobra.Command, args []string) {
	err := conf.LoadConfig(cmd)

	if err != nil {
		log.Fatal("Failed to load configuration: ", err.Error())
	}

}
