package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"	
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type FileInfo struct {
    Name    string
	Size    int64
	Dir		bool
    ModTime time.Time
}

type Room struct {
	Directory string
	Files     []FileInfo
}

func (r *Room) files() []string {
	files := []string{}

	for _, file := range r.Files {
		files = append(files, file.Name)
	}

	return files
}

var lsCmd = &cobra.Command{
	Use:   "ls [roomdir to print]",
	Short: "List roomdir contents",
	Long: `List roomdir contents`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get("http://localhost:3002/"+args[0])
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var room Room
		json.Unmarshal(body, &room)
		fmt.Println(room.files())
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
