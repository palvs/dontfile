package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [file to download]",
	Short: "Download a specific file",
	Long: `Download a specific file`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Create("test.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		res, err := http.Get("http://localhost:3002/"+args[0])
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		io.Copy(file, res.Body)

		fmt.Println("File successfully downloaded")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
