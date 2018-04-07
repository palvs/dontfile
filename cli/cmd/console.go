package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var baseApi string = "http://localhost:3002"
var roomDir string = ""
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
		main()
	},
}

func init() {
	rootCmd.AddCommand(consoleCmd)
}

func main() {
	var opt, arg string

	for opt != "exit" {	
		fmt.Printf("<Dontfile (%s)> ", baseApi + roomDir)
		switch fmt.Scanln(&opt, &arg); opt {
			case "ls": 
				ls()
			case "pwd":
				pwd()
			case "cd":
				cd(arg)
			case "rm":
				rm(arg)
			case "get":
				get(arg)
		}
	}
}

func get(file string) {
	fileUrl := baseApi + roomDir + "/" + file
	fileOut, _ := os.Create(file)
	defer fileOut.Close()

	res, _ := http.Get(fileUrl)
	defer res.Body.Close()

	io.Copy(fileOut, res.Body)
}

func ls() {
	res, _ := http.Get(baseApi + "/" + roomDir)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var room Room

	json.Unmarshal(body, &room)
	fmt.Println(room.files())
}

func cd(subdir string) {
	switch subdir {
		case "/":
			roomDir = "/"
		case ".":
		case "..":
			roomDir = strings.TrimSuffix(roomDir, "/" + strings.Split(roomDir, "/")[strings.Count(roomDir, "/")])
		default:
			roomDir += "/" + subdir
	}
}

func rm(file string) {
	fileUrl := baseApi + roomDir + "/" + file

	fmt.Println(fileUrl)

	req, _ := http.NewRequest("DELETE", fileUrl, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
}

func pwd() {
	fmt.Println(baseApi + "/"  + roomDir)
}