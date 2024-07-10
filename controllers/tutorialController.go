package controllers

import (
	"devwithgo/models"
	"devwithgo/util"
	"fmt"
	"html/template"
	"net/http"
	"strings"
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

func TutorialController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/tutorials/tutorial.html"))
	http.HandleFunc("/tutorial/", func(w http.ResponseWriter, r *http.Request) {

		urlPath := strings.TrimPrefix(r.URL.Path, "/tutorial/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get tutorial by URL
		getTutorial, err := models.TutorialFindByUrl(urlPath)
		if err != nil {
			fmt.Println("Error finding tutorial by URL:", err)
		}

		// Create raw content for html template
		tutorialRawContent := template.HTML(getTutorial.Content)

		data := TutorialData{
			PageTitle:          getTutorial.Title,
			PageDescription:    getTutorial.Description,
			CurrentYear:        time.Now().Year(),
			Tutorial:           getTutorial,
			TutorialContentRaw: tutorialRawContent,
		}

		// Redirect to 404 page if the content has been not published yet
		if !util.DateContentValidation(getTutorial.Published) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		tmpl.Execute(w, data)
	})
}
