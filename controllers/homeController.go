package controllers

import (
	"devwithgo/models"
	"html/template"
	"net/http"
)

func Home() {
	setPageData := models.NewPageData("Dev With Go: do it with Golang", "Go programming language: learn Go with guides, tutorials, and the latest news about Golang, the perfect resource for Gophers and aspiring Go developers")

	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := setPageData
		tmpl.Execute(w, data)
	})
}
