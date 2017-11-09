package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{[a-zA-Z0-9]+}", fileUpload)
	r.HandleFunc("/{[a-zA-Z0-9]+}/{[a-zA-Z0-9]+}", fileDownload)
	r.HandleFunc("/{[a-zA-Z0-9]+}/{[a-zA-Z0-9]+}/delete", fileDelete)
	r.HandleFunc("/", helloFunc)

	addr := ":" + os.Getenv("PORT")

	fmt.Println(addr)
	http.ListenAndServe(addr, r)
}

func helloFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Dontfile</title>
			<link rel="stylesheet"href="https://bootswatch.com/4/lumen/bootstrap.min.css">
		</head>
		<body>
		<div class="container">
		<h2>Use /{link} para compartilhar arquivos!</h2></body>
		</html>`)
}

func fileUpload(w http.ResponseWriter, req *http.Request) {

	cmd := exec.Command("mkdir", "storage")
	cmd.Run()

	dir := req.URL.Path[1:]

	if req.Method == http.MethodPost {
		file, header, err := req.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		fin := header.Filename

		fileBytes, err := ioutil.ReadAll(file)
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
	fmt.Fprintf(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Dontfile</title>
			<link rel="stylesheet"href="https://bootswatch.com/4/lumen/bootstrap.min.css">
		</head>
		<body>
		<br>
		<div class="container">
		<div class="row"><br>`)

	files, _ := ioutil.ReadDir("storage/" + dir)
	for _, file := range files {
		fmt.Fprintf(w, `
						<div class="col-md-3">
						    <div class="panel panel-default">
						        	<div class="panel-heading">
							          <p>%s</p>
							        </div>
							        <div class="panel-body">
							          <a class="btn btn-danger" href="%s">Deletar</a>
							          <a class="btn btn-primary" href="%s">Download</a>
						          </div>
						    </div>
						  
						</div>`, file.Name(), dir+"/"+file.Name()+"/delete", dir+"/"+file.Name())
	}

	fmt.Fprintf(w, `</div><hr><div class="row">
		<form class="form-horizontal" action="/%s" method="post" enctype="multipart/form-data">
						<fieldset>
				<div class="input-field">
					<div class="form-group">
						<input class="form-control" id="file" type="file" name="file" multiple>
					</div>
					<div class="form-group">
						<button class="btn btn-default btn-lg btn-block" type="submit">Enviar</button>
					</div>
				</div>
				<br>
				</fieldset>
			</form>
			</div>
			</div>
		</body>
		</html>`, dir)

}

func fileDownload(w http.ResponseWriter, req *http.Request) {
	dir := "storage/" + req.URL.Path[1:]
	http.ServeFile(w, req, dir)
}

func fileDelete(w http.ResponseWriter, req *http.Request) {
	dir := "storage/" + req.URL.Path[1:]
	dir = strings.TrimSuffix(dir, "/delete")

	cmd := exec.Command("rm", "-rf", dir)
	cmd.Run()

	dirs := strings.Split(dir, "/")

	http.Redirect(w, req, "/"+dirs[1], 301)
}
