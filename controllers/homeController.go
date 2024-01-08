package controllers

import (
	"html/template"
	"net/http"
)

type PageData struct {
	PageTitle string
}

func Home() {
	tmpl := template.Must(template.ParseFiles("./views/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle: "Dev With GO",
		}
		tmpl.Execute(w, data)
	})
}
