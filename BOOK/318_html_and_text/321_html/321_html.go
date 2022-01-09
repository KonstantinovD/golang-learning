package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var DATA []Entry
var tFile string

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	tmpl := template.Must(template.ParseFiles(tFile))
	tmpl.Execute(w, DATA)
}
func main() {
	tFile = "D:\\Danik_Prog\\Programming\\golang\\programs\\learning\\" +
		"BOOK\\318_html_and_text\\321_html\\gohtml.html"

	DATA = []Entry{
		{1, 2, 1},
		{2, 4, 4},
		{3, 6, 9},
		{11, 22, 121},
		{20, 40, 400},
	}

	http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
