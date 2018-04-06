package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"mime/multipart"
	"os"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push [roomdir] [file to upload]",
	Short: "Send a file",
	Long: `Send a file`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		data := &bytes.Buffer{}
		w := multipart.NewWriter(data)

		part, err := w.CreateFormFile("file", args[1])
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(part, file)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("POST", "http://localhost:3002/"+args[0], data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", w.FormDataContentType())
		
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		fmt.Println("File uploaded successfully")
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
