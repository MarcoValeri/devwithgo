package admincontrollers

import (
	"devwithgo/models"
	"devwithgo/util"
	"fmt"
	"html/template"
	"net/http"
)

type tutorialData struct {
	PageTitle                 string
	TitleError                string
	DescriptionError          string
	UrlError                  string
	PublishedError            string
	UpdatedError              string
	ImageError                string
	ContentError              string
	Images                    []models.Image
	Tutorials                 []models.Tutorial
	TutorialsWithRelatedImage []models.TutorialWithRelatedImage
}

func AdminTutorials() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-tutorials.html"))
	http.HandleFunc("/admin/tutorials", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["user-admin-authentication"] == true {
			tutorialsData, err := models.TutorialShowTutorials()
			if err != nil {
				fmt.Println("Error getting tutorialsData:", err)
			}

			data := tutorialData{
				PageTitle:                 "Tutorials",
				TutorialsWithRelatedImage: tutorialsData,
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}

func AdminTutorialAdd() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-tutorial-add.html"))
	http.HandleFunc("/admin/tutorial-add", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {

			imagesData, errImagesData := models.ImageShowImages()
			if errImagesData != nil {
				fmt.Println("Error getting imagesData:", errImagesData)
			}

			data := tutorialData{
				PageTitle: "Tutorial Add",
				Images:    imagesData,
			}

			// Flag validation
			var areAdminTutorialInputsValid [7]bool
			isFormSubmittionValid := false

			// Get values from the form
			getAdminTutorialTitle := r.FormValue("tutorial-title")
			getAdminTutorialDescription := r.FormValue("tutorial-description")
			getAdminTutorialUrl := r.FormValue("tutorial-url")
			getAdminTutorialPublished := r.FormValue("tutorial-published")
			getAdminTutorialUpdated := r.FormValue("tutorial-updated")
			getAdminTutorialImage := r.FormValue("tutorial-image")
			getAdminTutorialContent := r.FormValue("tutorial-content")
			getAdminTutorialAdd := r.FormValue("tutorial-add")

			// Sanitize form inputs
			getAdminTutorialTitle = util.FormSanitizeStringInput(getAdminTutorialTitle)
			getAdminTutorialDescription = util.FormSanitizeStringInput(getAdminTutorialDescription)
			getAdminTutorialUrl = util.FormSanitizeStringInput(getAdminTutorialUrl)
			getAdminTutorialPublished = util.FormSanitizeStringInput(getAdminTutorialPublished)
			getAdminTutorialUpdated = util.FormSanitizeStringInput(getAdminTutorialUpdated)
			getAdminTutorialImage = util.FormSanitizeStringInput(getAdminTutorialImage)
			getAdminTutorialAdd = util.FormSanitizeStringInput(getAdminTutorialAdd)

			// Check if the form has been submitted
			if getAdminTutorialAdd == "Add new guide" {
				// Title validation
				if len(getAdminTutorialTitle) > 0 {
					data.TitleError = ""
					areAdminTutorialInputsValid[0] = true
				} else {
					data.TitleError = "Title should be longer than 0"
					areAdminTutorialInputsValid[0] = false
				}

				// Description validation
				if len(getAdminTutorialDescription) > 0 {
					data.DescriptionError = ""
					areAdminTutorialInputsValid[1] = true
				} else {
					data.DescriptionError = "Description should be longer than 0"
					areAdminTutorialInputsValid[1] = false
				}

				// URL validation
				if len(getAdminTutorialUrl) > 0 {
					data.UrlError = ""
					areAdminTutorialInputsValid[2] = true
				} else {
					data.UrlError = "URL should be longer than 0"
					areAdminTutorialInputsValid[2] = false
				}

				// Published validation
				if len(getAdminTutorialPublished) > 0 {
					data.PublishedError = ""
					areAdminTutorialInputsValid[3] = true
				} else {
					data.PublishedError = "Add a date"
					areAdminTutorialInputsValid[3] = false
				}

				// Update validation
				if len(getAdminTutorialUpdated) > 0 {
					data.UpdatedError = ""
					areAdminTutorialInputsValid[4] = true
				} else {
					data.UpdatedError = "Add a date"
					areAdminTutorialInputsValid[4] = false
				}

				// Image validation
				if len(getAdminTutorialImage) > 0 {
					data.ImageError = ""
					areAdminTutorialInputsValid[5] = true
				} else {
					data.ImageError = "An image is required"
					areAdminTutorialInputsValid[5] = false
				}

				// Content validation
				if len(getAdminTutorialContent) > 0 {
					data.ContentError = ""
					areAdminTutorialInputsValid[6] = true
				} else {
					data.ContentError = "Content should be longher than 0"
					areAdminTutorialInputsValid[6] = false
				}

				for i := 0; i < len(areAdminTutorialInputsValid); i++ {
					isFormSubmittionValid = true
					if !areAdminTutorialInputsValid[i] {
						isFormSubmittionValid = false
						break
					}
				}

				// Create a new tutorial if all inputs are valid
				if isFormSubmittionValid {
					// Get image id for the relationship one-to-many between tutorials and images
					getAdminTutorialImageId, _ := models.ImageFindByUrlReturnItsId(getAdminTutorialImage)
					createNewTutorial := models.TutorialNew(
						1,
						getAdminTutorialTitle,
						getAdminTutorialDescription,
						getAdminTutorialUrl,
						getAdminTutorialPublished,
						getAdminTutorialUpdated,
						getAdminTutorialImageId,
						getAdminTutorialContent,
					)
					models.TutorialAddNewToDB(createNewTutorial)
					http.Redirect(w, r, "/admin/tutorials", http.StatusSeeOther)
				}
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
