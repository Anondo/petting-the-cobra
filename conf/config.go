package conf

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoadConfig(cmd *cobra.Command) error {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	// from the environment
	viper.SetEnvPrefix("NETLIFY")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	configFile, err := cmd.Flags().GetString("config")

	if err != nil {
		return err
	}

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		// from a config file
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/desktop/golang/src/cliapp")
	}

	// NOTE: this will require that you have config file somewhere in the paths specified. It can be reading from JSON, TOML, YAML, HCL, and Java properties files.
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	fmt.Printf("port: %d\n", viper.GetInt("port"))

	return nil

}