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
	PageTitle       string
	PageDescription string
	CurrentYear     int
	Guides          []models.Guide
	Guide           models.Guide
	GuideContentRaw template.HTML
}

func GuidesArchiveController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/guides/guides-archive.html"))
	http.HandleFunc("/guides/guides", func(w http.ResponseWriter, r *http.Request) {

		// Get all the guides
		getAllGuides, err := models.GuideGetPublishedGuides()
		if err != nil {
			fmt.Println("Error getting all guides:", err)
		}

		// Set data page
		data := GuideData{
			PageTitle:       "Learn Go programming language",
			PageDescription: "Learn Go programming language: guides from basics to advanced concepts, mastering powerful Golang features for becoming a skilled Gopher and Go developer",
			CurrentYear:     time.Now().Year(),
			Guides:          getAllGuides,
		}

		tmpl.Execute(w, data)
	})
}

func GuideController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/guides/guide.html"))
	http.HandleFunc("/guide/", func(w http.ResponseWriter, r *http.Request) {

		urlPath := strings.TrimPrefix(r.URL.Path, "/guide/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get guide by url
		getGuide, err := models.GuideFindByUrl(urlPath)
		if err != nil {
			fmt.Println("Error finding guide by url:", err)
		}

		// Create raw content for html template
		guideContentRaw := template.HTML(getGuide.Content)

		// Set data page
		data := GuideData{
			PageTitle:       getGuide.Title,
			PageDescription: getGuide.Description,
			CurrentYear:     time.Now().Year(),
			Guide:           getGuide,
			GuideContentRaw: guideContentRaw,
		}

		// Redirect to 404 page if the content has been not published yet
		if !util.DateContentValidation(getGuide.Published) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		tmpl.Execute(w, data)
	})
}
