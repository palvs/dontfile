package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [file to delete]",
	Short: "Delete a file",
	Long: `Delete a file`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest("DELETE", "http://localhost:3002/"+args[0], nil)
		if err != nil {
			log.Fatal(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		fmt.Println("File successfully deleted")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
