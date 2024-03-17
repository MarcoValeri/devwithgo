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

type GuideData struct {
	PageTitle   string
	CurrentYear int
	Guides      []models.Guide
	Guide       models.Guide
}

func GuidesArchive() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/guides/guides-archive.html"))
	http.HandleFunc("/guides/guides-all-content", func(w http.ResponseWriter, r *http.Request) {

		// Get all the guides
		getAllGuides, err := models.GuideShowGuides()
		if err != nil {
			fmt.Println("Error getting all guides:", err)
		}

		// Set data page
		data := GuideData{
			PageTitle:   "Go Guides",
			CurrentYear: time.Now().Year(),
			Guides:      getAllGuides,
		}

		tmpl.Execute(w, data)
	})
}

func Guide() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/guides/guide.html"))
	http.HandleFunc("/guide/", func(w http.ResponseWriter, r *http.Request) {

		urlPath := strings.TrimPrefix(r.URL.Path, "/guide/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get guide by url
		getGuide, err := models.GuideFindByUrl(urlPath)
		if err != nil {
			fmt.Println("Error finding guide by url:", err)
		}

		// Set data page
		data := GuideData{
			PageTitle:   "Go Guides",
			CurrentYear: time.Now().Year(),
			Guide:       getGuide,
		}

		tmpl.Execute(w, data)
	})
}
