package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type EchoRes struct {
	Result string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err = template.ParseFiles("templates/index.html")
	if r.URL.Path != "/" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, "")
}
func EchoPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/echo" {
		http.Error(w, "Path not found.", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	input := r.PostFormValue("text")
	output := EchoRes{Result: input}
	fmt.Fprintln(w, output)
}

var tmpl *template.Template
var err error

func main() {
	fmt.Println("server is working normally on port :8080, press ctrl + c to terminate server")
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/echo", EchoPage)
	http.ListenAndServe(":8080", mux)
}
