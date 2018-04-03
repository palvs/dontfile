package main

import (
	"fmt"
	//"strings"
	"net/http"
	"io/ioutil"
	"io"
	"net/url"
	"os"
)

var base string = "http://localhost:3001/"
var dir string = "test"

func main() {
	
	upload("test.txt")
}

func download(file string) {
	fileUrl := base + dir + file

	fileOut, _ := os.Create(file)
	defer fileOut.Close()

	res, _ := http.Get(fileUrl)
	defer res.Body.Close()

	io.Copy(fileOut, res.Body)
}

func upload(file string) {
	fileUpload, _ := os.Open(file)
	defer fileUpload.Close()

	//res, _ := http.Post(base + dir, "binary/octet-stream", fileUpload)
	res, _ := http.PostForm(base + dir, url.Values{"file": fileUpload})
	defer res.Body.Close()

	fmt.Println("Ok")
}

func ls() {
	res, _ := http.Get(base + dir)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func cd(subdir string) {
	dir += "/" + subdir
}

func pwd() {
	fmt.Println(dir)
}