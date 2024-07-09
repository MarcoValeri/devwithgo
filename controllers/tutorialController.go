package controllers

import (
	"devwithgo/models"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type TutorialData struct {
	PageTitle          string
	PageDescription    string
	CurrentYear        int
	Tutorials          []models.TutorialWithRelatedImage
	Tutorial           models.TutorialWithRelatedImage
	TutorialContentRaw template.HTML
}

func TutorialsArchiveController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/tutorials/tutorials-archive.html"))
	http.HandleFunc("/tutorials/tutorials", func(w http.ResponseWriter, r *http.Request) {
		// Get all the tutorials
		getAllTutorials, err := models.TutorialGetPublishedTutorials()
		if err != nil {
			fmt.Println("Error getting all tutorials:", err)
		}

		// Set data page
		data := TutorialData{
			PageTitle:       "Go tutorials",
			PageDescription: "Go tutorials",
			CurrentYear:     time.Now().Year(),
			Tutorials:       getAllTutorials,
		}

		tmpl.Execute(w, data)
	})
}
