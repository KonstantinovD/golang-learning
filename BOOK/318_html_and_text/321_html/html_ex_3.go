package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData2 struct {
	Title string
	Users []string
}

func main() {

	filePath := "D:\\Danik_Prog\\Programming\\golang\\programs\\learning\\" +
		"BOOK\\318_html_and_text\\321_html\\index.html"

	data := ViewData2{
		Title: "Users List",
		Users: []string{"Tom", "Bob", "Sam"},
	}
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			tmpl, _ := template.ParseFiles(filePath)
			tmpl.Execute(w, data)
		})
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
