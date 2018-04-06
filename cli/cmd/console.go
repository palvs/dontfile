package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Dontfile console",
	Long: `
 _____              _    __ _ _                 _ _ 
|  __ \            | |  / _(_) |               | (_)
| |  | | ___  _ __ | |_| |_ _| | ___ ______ ___| |_ 
| |  | |/ _ \| '_ \| __|  _| | |/ _ \______/ __| | |
| |__| | (_) | | | | |_| | | | |  __/     | (__| | |
|_____/ \___/|_| |_|\__|_| |_|_|\___|      \___|_|_|    
	
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("console called")
	},
}

func init() {
	rootCmd.AddCommand(consoleCmd)
}
