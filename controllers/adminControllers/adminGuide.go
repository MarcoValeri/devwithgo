package admincontrollers

import (
	"devwithgo/models"
	"devwithgo/util"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type guideData struct {
	PageTitle        string
	TitleError       string
	DescriptionError string
	UrlError         string
	PublishedError   string
	UpdatedError     string
	ContentError     string
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
			data := guideData{
				PageTitle: "Guide Add",
			}

			// Flag validation
			var areAdminGuideInputsValid [6]bool
			isFormSubmittionValid := false

			// Get values from the form
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
				if len(getAdminGuideTitle) > 0 {
					data.TitleError = ""
					areAdminGuideInputsValid[0] = true
				} else {
					data.TitleError = "Title should be longher than 0"
					areAdminGuideInputsValid[0] = false
				}

				// Description validation
				if len(getAdminGuideDescription) > 0 && len(getAdminGuideDescription) < 200 {
					data.DescriptionError = ""
					areAdminGuideInputsValid[1] = true
				} else {
					data.DescriptionError = "Description should be between 1 to 160 charactes"
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
					data.UpdatedError = ""
					areAdminGuideInputsValid[4] = true
				} else {
					data.UpdatedError = "Add a data"
					areAdminGuideInputsValid[4] = false
				}

				// Content validation
				if len(getAdminGuideContent) > 0 {
					data.ContentError = ""
					areAdminGuideInputsValid[5] = true
				} else {
					data.ContentError = "Add the content"
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

func AdminGuideEdit() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-guide-edit.html"))
	http.HandleFunc("/admin/guide-edit/", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["user-admin-authentication"] == true {
			idPath := strings.TrimPrefix(r.URL.Path, "/admin/guide-edit/")
			idPath = util.FormSanitizeStringInput(idPath)

			guideId, err := strconv.Atoi(idPath)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}

			getGuideEdit, err := models.GuideFindIt(guideId)
			if err != nil {
				fmt.Println("Error to find user:", err)
			}

			// Create data for the page
			data := guideData{
				PageTitle:  "Admin Guide Edit",
				TitleError: "",
				Guides:     getGuideEdit,
			}

			/**
			* Check if the form for editing the guide has been submitted
			* and
			* validate the inputs
			 */
			var areAdmingGuideEditInputsValid [6]bool
			isFormSubmittionValid := false

			// Get the values from the form
			getAdminGuideTitleEdit := r.FormValue("guide-title-edit")
			getAdminGuideDescriptionEdit := r.FormValue("guide-description-edit")
			getAdminGuideUrlEdit := r.FormValue("guide-url-edit")
			getAdminGuidePublishedEdit := r.FormValue("guide-published-edit")
			getAdminGuideUpdatedEdit := r.FormValue("guide-updated-edit")
			getAdminGuideContentEdit := r.FormValue("guide-content-edit")
			getAdminGuideSubmitEdit := r.FormValue("guide-edit")

			// Sanitize form inputs
			getAdminGuideTitleEdit = util.FormSanitizeStringInput(getAdminGuideTitleEdit)
			getAdminGuideDescriptionEdit = util.FormSanitizeStringInput(getAdminGuideDescriptionEdit)
			getAdminGuideUrlEdit = util.FormSanitizeStringInput(getAdminGuideUrlEdit)
			getAdminGuidePublishedEdit = util.FormSanitizeStringInput(getAdminGuidePublishedEdit)
			getAdminGuideUpdatedEdit = util.FormSanitizeStringInput(getAdminGuideUpdatedEdit)
			getAdminGuideSubmitEdit = util.FormSanitizeStringInput(getAdminGuideSubmitEdit)

			// Check if the form has been submitted
			if getAdminGuideSubmitEdit == "Edit this guide" {
				// Title validation
				if len(getAdminGuideTitleEdit) > 0 && len(getAdminGuideTitleEdit) < 60 {
					data.TitleError = ""
					areAdmingGuideEditInputsValid[0] = true
				} else {
					data.TitleError = "Title should be between 1 to 60 charactes"
					areAdmingGuideEditInputsValid[0] = false
				}

				// Description validation
				if len(getAdminGuideDescriptionEdit) > 0 && len(getAdminGuideDescriptionEdit) < 200 {
					data.DescriptionError = ""
					areAdmingGuideEditInputsValid[1] = true
				} else {
					data.DescriptionError = "Description should be between 1 to 200 charactes"
					areAdmingGuideEditInputsValid[1] = false
				}

				// Url validation
				if len(getAdminGuideUrlEdit) > 0 {
					data.UrlError = ""
					areAdmingGuideEditInputsValid[2] = true
				} else {
					data.UrlError = "Add an URL"
					areAdmingGuideEditInputsValid[2] = false
				}

				// Published validation
				if len(getAdminGuidePublishedEdit) > 0 {
					data.PublishedError = ""
					areAdmingGuideEditInputsValid[3] = true
				} else {
					data.PublishedError = "Add a data"
					areAdmingGuideEditInputsValid[3] = false
				}

				// Updated validation
				if len(getAdminGuideUpdatedEdit) > 0 {
					data.UpdatedError = ""
					areAdmingGuideEditInputsValid[4] = true
				} else {
					data.UpdatedError = "Add a data"
					areAdmingGuideEditInputsValid[4] = false
				}

				// Content validation
				if len(getAdminGuideContentEdit) > 0 {
					data.ContentError = ""
					areAdmingGuideEditInputsValid[5] = true
				} else {
					data.ContentError = "Add the content"
					areAdmingGuideEditInputsValid[5] = false
				}

				for i := 0; i < len(areAdmingGuideEditInputsValid); i++ {
					isFormSubmittionValid = true
					if !areAdmingGuideEditInputsValid[i] {
						isFormSubmittionValid = false
						break
					}
				}

				// Edit guide if all inputs are valid and redirect to all guides list
				if isFormSubmittionValid {
					edtiGuide := models.GuideNew(guideId, getAdminGuideTitleEdit, getAdminGuideDescriptionEdit, getAdminGuideUrlEdit, getAdminGuidePublishedEdit, getAdminGuideUpdatedEdit, getAdminGuideContentEdit)
					models.GuideEdit(edtiGuide)
					http.Redirect(w, r, "/admin/guides", http.StatusSeeOther)
				}
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}

func AdminGuideDelete() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-guide-delete.html"))
	http.HandleFunc("/admin/guide-delete/", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {
			idPath := strings.TrimPrefix(r.URL.Path, "/admin/guide-delete/")
			idPath = util.FormSanitizeStringInput(idPath)

			guideId, err := strconv.Atoi(idPath)
			if err != nil {
				fmt.Println("Error converting string to integert:", err)
				return
			}

			getGuideDelete, err := models.GuideFindIt(guideId)
			if err != nil {
				fmt.Println("Error to find the user:", err)
			}

			data := guideData{
				PageTitle: "Admin Delete Guide",
				Guides:    getGuideDelete,
			}

			/**
			* Check if the form for deleting user has
			* been submitted
			* and
			* delete the selected user
			 */
			isFormSubmittionValid := false

			// Get the value from the form
			getAdminGuideDeleteSubmit := r.FormValue("admin-guide-delete")

			// Sanitize form input
			getAdminGuideDeleteSubmit = util.FormSanitizeStringInput(getAdminGuideDeleteSubmit)

			// Check if the form has been submitted
			if getAdminGuideDeleteSubmit == "Delete this guide" {
				isFormSubmittionValid = true
			}

			if isFormSubmittionValid {
				models.GuideDelete(guideId)
				http.Redirect(w, r, "/admin/guides", http.StatusSeeOther)
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}
