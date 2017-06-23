package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{[a-zA-Z0-9]+}", fileUpload)
	r.HandleFunc("/{[a-zA-Z0-9]+}/{[a-zA-Z0-9]+}", fileDownload)
	r.HandleFunc("/", helloFunc)

	addr := ":"+os.Getenv("PORT")

	fmt.Println("Ouvindo em", addr)
	http.ListenAndServe(addr, r)
}

func helloFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Dontfile</title>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.99.0/css/materialize.min.css">
		</head>
		<body>
		<div class="container">
		<h2>Use /{link} para compartilhar arquivos!</h2></body>
		</html>`)
}

func fileUpload(w http.ResponseWriter, req *http.Request) {

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

		cmd := exec.Command("mkdir", dir)
		cmd.Run()

		err = ioutil.WriteFile(dir+"/"+fin, fileBytes, 0644)
		if err != nil {
			panic(err)
		}
	}
	fmt.Fprintf(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Dontfile</title>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.99.0/css/materialize.min.css">
		</head>
		<body>
		<div class="container">`)

	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		//fmt.Fprintf(w, `<a href="%s">%s</a>`, dir+"/"+file.Name(), file.Name())
		fmt.Fprintf(w, `<div class="col s12 m7">
						    <div class="card horizontal">
						      <div class="card-stacked">
						        <div class="card-content">
						          <p>%s</p>
						        </div>
						        <div class="card-action">
						          <a href="%s">Download</a>
						        </div>
						      </div>
						    </div>
						  </div>`, file.Name(), dir+"/"+file.Name())
	}
	//fmt.Fprintf(w, )

	fmt.Fprintf(w, `<form action="/%s" method="post" enctype="multipart/form-data">
				<div class="input-field col s6">
					<input id="file" type="file" name="file">
					<button class="btn waves-effect waves-light" type="submit">Enviar</button>
				</div>
				<br>
				
			</form>
			</div>
		</body>
		</html>`, dir)

}

func fileDownload(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, req.URL.Path[1:])
}
