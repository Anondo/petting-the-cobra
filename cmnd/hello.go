package cmnd

import "github.com/spf13/cobra"

func Hello() *cobra.Command {
	return &cobra.Command{
		Use:   "greet",
		Short: "Say hello basically!!",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Hello " + args[0])
			return nil
		},
	}
}
