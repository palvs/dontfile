package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dontfile",
	Short: "Dontfile-cli",
	Long: `
 _____              _    __ _ _                 _ _ 
|  __ \            | |  / _(_) |               | (_)
| |  | | ___  _ __ | |_| |_ _| | ___ ______ ___| |_ 
| |  | |/ _ \| '_ \| __|  _| | |/ _ \______/ __| | |
| |__| | (_) | | | | |_| | | | |  __/     | (__| | |
|_____/ \___/|_| |_|\__|_| |_|_|\___|      \___|_|_|    
	
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".test")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
