package admincontrollers

import (
	"devwithgo/models"
	"devwithgo/util"
	"fmt"
	"html/template"
	"net/http"
)

type guideData struct {
	PageTitle        string
	TitleError       string
	DescriptionError string
	UrlError         string
	PublishedError   string
	UpdatedError     string
	ContentErorr     string
	Guides           []models.Guide
}

func AdminGuides() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-guides.html"))
	http.HandleFunc("/admin/guides", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {
			guidesData, err := models.GuideShowGuides()
			if err != nil {
				fmt.Println("Error getting guidesData: %w", err)
			}

			data := guideData{
				PageTitle: "Guides",
				Guides:    guidesData,
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}

func AdminGuideAdd() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-guide-add.html"))
	http.HandleFunc("/admin/guide-add", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {
			// TODO: get and show all guides
			data := guideData{
				PageTitle: "Guide Add",
			}

			// Flag validation
			var areAdminGuideInputsValid [6]bool
			isFormSubmittionValid := false

			// Get value from the form
			getAdminGuideTitle := r.FormValue("guide-title")
			getAdminGuideDescription := r.FormValue("guide-description")
			getAdminGuideUrl := r.FormValue("guide-url")
			getAdminGuidePublished := r.FormValue("guide-published")
			getAdminGuideUpdated := r.FormValue("guide-updated")
			getAdminGuideContent := r.FormValue("guide-content")
			getAdminGuideAdd := r.FormValue("guide-add")

			// Sanitize form inputs
			getAdminGuideTitle = util.FormSanitizeStringInput(getAdminGuideTitle)
			getAdminGuideDescription = util.FormSanitizeStringInput(getAdminGuideDescription)
			getAdminGuideUrl = util.FormSanitizeStringInput(getAdminGuideUrl)
			getAdminGuidePublished = util.FormSanitizeStringInput(getAdminGuidePublished)
			getAdminGuideUpdated = util.FormSanitizeStringInput(getAdminGuideUpdated)
			getAdminGuideAdd = util.FormSanitizeStringInput(getAdminGuideAdd)

			// Check if the form has been submitted
			if getAdminGuideAdd == "Add new guide" {
				// Title validation
				if len(getAdminGuideTitle) > 0 && len(getAdminGuideTitle) < 60 {
					data.TitleError = ""
					areAdminGuideInputsValid[0] = true
				} else {
					data.TitleError = "Title should be between 1 to 60 charactes"
					areAdminGuideInputsValid[0] = false
				}

				// Description validation
				if len(getAdminGuideDescription) > 0 && len(getAdminGuideDescription) < 160 {
					data.DescriptionError = ""
					areAdminGuideInputsValid[1] = true
				} else {
					data.DescriptionError = "Description should be between 1 t0 160 charactes"
					areAdminGuideInputsValid[1] = false
				}

				// Url validation
				if len(getAdminGuideUrl) > 0 {
					data.UrlError = ""
					areAdminGuideInputsValid[2] = true
				} else {
					data.UrlError = "Add an URL"
					areAdminGuideInputsValid[2] = false
				}

				// Published validation
				if len(getAdminGuidePublished) > 0 {
					data.PublishedError = ""
					areAdminGuideInputsValid[3] = true
				} else {
					data.PublishedError = "Add a data"
					areAdminGuideInputsValid[3] = false
				}

				// Updated validation
				if len(getAdminGuideUpdated) > 0 {
					data.PublishedError = ""
					areAdminGuideInputsValid[4] = true
				} else {
					data.UpdatedError = "Add a data"
					areAdminGuideInputsValid[4] = false
				}

				// Content validation
				if len(getAdminGuideContent) > 0 {
					data.ContentErorr = ""
					areAdminGuideInputsValid[5] = true
				} else {
					data.UpdatedError = "Add a data"
					areAdminGuideInputsValid[5] = false
				}

				for i := 0; i < len(areAdminGuideInputsValid); i++ {
					isFormSubmittionValid = true
					if !areAdminGuideInputsValid[i] {
						isFormSubmittionValid = false
						break
					}
				}

				// Create a new guide if all inputs are valid
				if isFormSubmittionValid {
					createNewGuide := models.GuideNew(
						1,
						getAdminGuideTitle,
						getAdminGuideDescription,
						getAdminGuideUrl,
						getAdminGuidePublished,
						getAdminGuideUpdated,
						getAdminGuideContent,
					)
					models.GuideAddNewToDB(createNewGuide)
					http.Redirect(w, r, "/admin/guides", http.StatusSeeOther)
				}
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}
