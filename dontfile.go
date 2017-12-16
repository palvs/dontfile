package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Upload struct {
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
		upload, header, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer upload.Close()

		fin := header.Filename

		fileBytes, err := ioutil.ReadAll(upload)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("mkdir", "storage/"+dir)
		cmd.Run()

		err = ioutil.WriteFile("storage/"+dir+"/"+fin, fileBytes, 0644)
		if err != nil {
			panic(err)
		}
	}

	files, _ := ioutil.ReadDir("storage/" + dir)

	u := Upload{dir, files}

	tt, _ := template.ParseFiles("views/upload.html")
	tt.Execute(w, u)

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
