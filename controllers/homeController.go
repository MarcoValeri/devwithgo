package controllers

import (
	"devwithgo/models"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type HomeData struct {
	PageTitle       string
	PageDescription string
	CurrentYear     int
	Tutorials       []models.TutorialWithRelatedImage
}

func Home() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Get last three published tutorials
		getLastThreePublishedTutorials, err := models.TutorialGetLimitPublishedTutorials(3)
		if err != nil {
			fmt.Println("Error getting last three tutorials:", err)
		}

		// Set date page
		data := HomeData{
			PageTitle:       "Dev With Go: do it with Golang",
			PageDescription: "Go programming language: learn Go with guides, tutorials, and the latest news about Golang, the perfect resource for Gophers and aspiring Go developers",
			CurrentYear:     time.Now().Year(),
			Tutorials:       getLastThreePublishedTutorials,
		}

		tmpl.Execute(w, data)
	})
}
