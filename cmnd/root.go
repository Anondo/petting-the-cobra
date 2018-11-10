package cmnd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Rootcmd = &cobra.Command{
		Use:   "greeter",
		Short: "greeter is a short application to greet people",
	}
)

func init() {
	Rootcmd.AddCommand(Hello())
}

func Execute() {
	if err := Rootcmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
