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

	configFile := viper.GetString("config")
	viper.BindEnv("name")

	name := viper.GetString("name")
	if name == "" {
		fmt.Println("Warning: name not provided")
	} else {
		fmt.Println(name)
	}

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		// from a config file
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
		//viper.AddConfigPath("$HOME/desktop/golang/src/petting-the-cobra")
	}

	// NOTE: this will require that you have config file somewhere in the paths specified. It can be reading from JSON, TOML, YAML, HCL, and Java properties files.
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	port := viper.GetInt("port")
	if port == 0 {
		fmt.Printf("port: %d\n", viper.GetInt("settings.port"))
	} else {
		fmt.Printf("port: %d\n", viper.GetInt("port"))
	}

	return nil

}
