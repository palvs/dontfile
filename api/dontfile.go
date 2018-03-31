package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gorilla/mux"
)


type FileInfo struct {
    Name    string
    Size    int64
    ModTime time.Time
}

type Room struct {
	Directory string
	Files     []FileInfo
}

var STORAGE_DIR = "../storage/"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{[a-zA-Z0-9]+}", fileUpload)
	r.HandleFunc("/{[a-zA-Z0-9]+}/{[a-zA-Z0-9]+}", fileDownload)
	r.HandleFunc("/{[a-zA-Z0-9]+}/{[a-zA-Z0-9]+}/delete", fileDelete)

	addr := ":" + os.Getenv("PORT")

	fmt.Println(addr)
	http.ListenAndServe(addr, r)
}

func fileUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cmd := exec.Command("mkdir", STORAGE_DIR)
	cmd.Run()

	dir := r.URL.Path[1:]

	if r.Method == http.MethodPost {
		_ = r.ParseMultipartForm(100000)
		m := r.MultipartForm

		files := m.File["file"]

		for i := range files {
			file, _ := files[i].Open()
			defer file.Close()

			cmd := exec.Command("mkdir", STORAGE_DIR + dir)
			cmd.Run()

			destinationFile, _ := os.Create(STORAGE_DIR + dir + "/" + files[i].Filename)
			defer destinationFile.Close()

			io.Copy(destinationFile, file)
		}
	}

	raw_files, _ := ioutil.ReadDir(STORAGE_DIR + dir)
	var files []FileInfo 

	for _, raw_file := range raw_files { 
		files = append(files, FileInfo{
			Name:	 raw_file.Name(),
			Size:	 raw_file.Size(),
			ModTime: raw_file.ModTime(),
		}) 
	}

	room := Room{dir, files}

	json, _ := json.Marshal(room)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func fileDownload(w http.ResponseWriter, r *http.Request) {
	dir := STORAGE_DIR + r.URL.Path[1:]
	http.ServeFile(w, r, dir)
}

func fileDelete(w http.ResponseWriter, r *http.Request) {
	dir := STORAGE_DIR + r.URL.Path[1:]
	dir = strings.TrimSuffix(dir, "/delete")

	cmd := exec.Command("rm", "-rf", dir)
	cmd.Run()

	dirs := strings.Split(dir, "/")

	http.Redirect(w, r, "/"+dirs[1], 301)
}
