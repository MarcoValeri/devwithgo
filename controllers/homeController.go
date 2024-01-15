package controllers

import (
	"devwithgo/models"
	"html/template"
	"net/http"
)

func Home() {
	setPageData := models.NewPageData("Do it with GO")

	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := setPageData
		tmpl.Execute(w, data)
	})
}
