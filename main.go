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
	if r.URL.Path != "/" {
		http.Error(w, "404", http.StatusNotFound)
		fmt.Fprintln(w, nil)
	}
	if r.Method != "GET" {
		http.Error(w, "404", http.StatusNotFound)
		fmt.Fprintln(w, "")
	}
}
func EchoPage(w http.ResponseWriter, r *http.Request) {
	input := r.PostFormValue("text")
	output := EchoRes{Result: input}
	fmt.Fprintln(w, output)
}

var tmpl *template.Template
var err error

func main() {
	fmt.Println("server is working normally")
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/echo", EchoPage)
}
