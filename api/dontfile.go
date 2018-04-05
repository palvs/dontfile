package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	r.HandleFunc("/{room:[^.]+[^./]}", fileIndex).Methods("GET")
	r.HandleFunc("/{room:[^.]+[^./]}", fileUpload).Methods("POST")
	r.HandleFunc("/{room:[^.]+[^./]}{file}", fileDownload).Methods("GET")
	r.HandleFunc("/{room:[^.]+[^./]}{file}", fileDelete).Methods("DELETE")

	addr := ":" + os.Getenv("PORT")

	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(addr, handler)
}

func fileIndex(w http.ResponseWriter, r *http.Request) {
	dir := mux.Vars(r)["room"]

	rawFiles, _ := ioutil.ReadDir(STORAGE_DIR + dir)
	var files []FileInfo 

	for _, rawFile := range rawFiles { 
		files = append(files, FileInfo{
			Name:	 rawFile.Name(),
			Size:	 rawFile.Size(),
			ModTime: rawFile.ModTime(),
		}) 
	}

	room := Room{dir, files}

	json, _ := json.Marshal(room)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func fileUpload(w http.ResponseWriter, r *http.Request) {
	dir := mux.Vars(r)["room"]

	cmd := exec.Command("mkdir", "-p", STORAGE_DIR + dir)
	cmd.Run()

	_ = r.ParseMultipartForm(100000)
	m := r.MultipartForm

	files := m.File["file"]

	for i := range files {
		file, _ := files[i].Open()
		defer file.Close()

		destinationFile, _ := os.Create(STORAGE_DIR + dir + "/" + files[i].Filename)
		defer destinationFile.Close()

		io.Copy(destinationFile, file)
	}

}

func fileDownload(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	file := v["room"] + v["file"]

	dir := STORAGE_DIR + file
	http.ServeFile(w, r, dir)
}

func fileDelete(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	file := v["room"] + v["file"]
	dir := STORAGE_DIR + file

	cmd := exec.Command("rm", "-rf", dir)
	cmd.Run()
}
