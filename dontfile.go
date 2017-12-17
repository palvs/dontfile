package main

import (
	//"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

type Room struct {
	Directory string
	Files     []os.FileInfo
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)

	r.HandleFunc("/{[a-zA-Z0-9]+}", fileUpload)
	r.HandleFunc("/{[a-zA-Z0-9]+}/{[a-zA-Z0-9]+}", fileDownload)
	r.HandleFunc("/{[a-zA-Z0-9]+}/{[a-zA-Z0-9]+}/delete", fileDelete)

	addr := ":" + os.Getenv("PORT")

	fmt.Println(addr)
	http.ListenAndServe(addr, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, nil)
}

func fileUpload(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("mkdir", "storage")
	cmd.Run()

	dir := r.URL.Path[1:]

	if r.Method == http.MethodPost {
		_ = r.ParseMultipartForm(100000)
		m := r.MultipartForm

		files := m.File["file"]

		for i := range files {
			file, _ := files[i].Open()
			defer file.Close()

			cmd := exec.Command("mkdir", "storage/"+dir)
			cmd.Run()

			destinationFile, _ := os.Create("storage/" + dir + "/" + files[i].Filename)
			defer destinationFile.Close()

			io.Copy(destinationFile, file)
		}
	}

	files, _ := ioutil.ReadDir("storage/" + dir)

	room := Room{dir, files}

	t, _ := template.ParseFiles("views/upload.html")
	t.Execute(w, room)

	/*
		json, _ := json.Marshal(room)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	*/

}

func fileDownload(w http.ResponseWriter, r *http.Request) {
	dir := "storage/" + r.URL.Path[1:]
	http.ServeFile(w, r, dir)
}

func fileDelete(w http.ResponseWriter, r *http.Request) {
	dir := "storage/" + r.URL.Path[1:]
	dir = strings.TrimSuffix(dir, "/delete")

	cmd := exec.Command("rm", "-rf", dir)
	cmd.Run()

	dirs := strings.Split(dir, "/")

	http.Redirect(w, r, "/"+dirs[1], 301)
}
