package controllers

import (
	"html/template"
	"net/http"
	"time"
)

type PageData struct {
	PageTitle   string
	CurrentYear int
}

func Home() {
	setCurrentYear := time.Now().Year()

	tmpl := template.Must(template.ParseFiles("./views/home.html", "./views/templates/head.html", "./views/templates/header.html", "./views/templates/footer.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle:   "Do it with GO",
			CurrentYear: setCurrentYear,
		}
		tmpl.Execute(w, data)
	})
}
