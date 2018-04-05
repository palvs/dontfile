package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"io"
	//"net/url"
	"os"
	"time"
	"encoding/json"
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

var base string = "http://localhost:3002"
var dir string = ""

func main() {
	var opt, arg string

	for opt != "exit" {	
		fmt.Printf("<Dontfile (%s)> ", base + dir)
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
	fileUrl := base + dir + "/" + file
	fileOut, _ := os.Create(file)
	defer fileOut.Close()

	res, _ := http.Get(fileUrl)
	defer res.Body.Close()

	io.Copy(fileOut, res.Body)
}

/*
func upload(file string) {
	fileUpload, _ := os.Open(file)
	defer fileUpload.Close()

	//res, _ := http.Post(base + dir, "binary/octet-stream", fileUpload)
	res, _ := http.PostForm(base + dir, url.Values{"file": fileUpload})
	defer res.Body.Close()

	fmt.Println("Ok")
}
*/

func ls() {
	res, _ := http.Get(base + "/" + dir)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var room Room

	json.Unmarshal(body, &room)
	fmt.Println(room.files())
}

func cd(subdir string) {
	switch subdir {
		case "/":
			dir = "/"
		case ".":
		case "..":
			dir = strings.TrimSuffix(dir, "/" + strings.Split(dir, "/")[strings.Count(dir, "/")])
		default:
			dir += "/" + subdir
	}
}

func rm(file string) {
	fileUrl := base + dir + "/" + file

	fmt.Println(fileUrl)

	req, _ := http.NewRequest("DELETE", fileUrl, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
}

func pwd() {
	fmt.Println(base + "/"  + dir)
}