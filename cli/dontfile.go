package main

import (
	"fmt"
	//"strings"
	"net/http"
	"io/ioutil"
	"os"
)

var dir string = "http://localhost:3001/test"

func main() {


	fmt.Println("cli")
	fmt.Println(os.Args[1])

	ls()
}

func download(file string) {
	file_url := dir + file

	fmt.Println(file_url)
 
	req, _ := http.NewRequest("GET", file_url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func upload() {

}

func ls() {
	req, _ := http.NewRequest("GET", dir, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func cd(subdir string) {
	
}

func pwd() {

}