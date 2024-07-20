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
			if getAdminTutorialAdd == "Add new tutorial" {
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

func AdminTutorialEdit() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-tutorial-edit.html"))
	http.HandleFunc("/admin/tutorial-edit/", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["user-admin-authentication"] == true {
			idPath := strings.TrimPrefix(r.URL.Path, "/admin/tutorial-edit/")
			idPath = util.FormSanitizeStringInput(idPath)

			tutorialId, err := strconv.Atoi(idPath)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}

			getTutorialEdit, err := models.TutorialWithRelatedImageFindById(tutorialId)
			if err != nil {
				fmt.Println("Error to find tutorial:", err)
				return
			}

			imagesData, errImagesData := models.ImageShowImages()
			if errImagesData != nil {
				fmt.Println("Error getting imagesData:", errImagesData)
			}

			// Create data for the page
			data := tutorialData{
				PageTitle:                 "Admin Tutorial Edit",
				TutorialsWithRelatedImage: getTutorialEdit,
				Images:                    imagesData,
			}

			/**
			* Check if the form for editing the tutorial has been submitted
			* and
			* validate the inputs
			 */
			var areAdminTutorialEditInputsValid [7]bool
			isFormSubmittionValid := false

			// Get the values from the inputs
			// Get values from the form
			getAdminTutorialTitleEdit := r.FormValue("tutorial-edit-title")
			getAdminTutorialDescriptionEdit := r.FormValue("tutorial-edit-description")
			getAdminTutorialUrlEdit := r.FormValue("tutorial-edit-url")
			getAdminTutorialPublishedEdit := r.FormValue("tutorial-edit-published")
			getAdminTutorialUpdatedEdit := r.FormValue("tutorial-edit-updated")
			getAdminTutorialImageEdit := r.FormValue("tutorial-edit-image")
			getAdminTutorialContentEdit := r.FormValue("tutorial-edit-content")
			getAdminTutorialEdit := r.FormValue("tutorial-edit")

			// Sanitize form inputs
			getAdminTutorialTitleEdit = util.FormSanitizeStringInput(getAdminTutorialTitleEdit)
			getAdminTutorialDescriptionEdit = util.FormSanitizeStringInput(getAdminTutorialDescriptionEdit)
			getAdminTutorialUrlEdit = util.FormSanitizeStringInput(getAdminTutorialUrlEdit)
			getAdminTutorialPublishedEdit = util.FormSanitizeStringInput(getAdminTutorialPublishedEdit)
			getAdminTutorialUpdatedEdit = util.FormSanitizeStringInput(getAdminTutorialUpdatedEdit)
			getAdminTutorialImageEdit = util.FormSanitizeStringInput(getAdminTutorialImageEdit)
			getAdminTutorialEdit = util.FormSanitizeStringInput(getAdminTutorialEdit)

			if getAdminTutorialEdit == "Edit this tutorial" {
				// Title validation
				if len(getAdminTutorialTitleEdit) > 0 {
					data.TitleError = ""
					areAdminTutorialEditInputsValid[0] = true
				} else {
					data.TitleError = "Title should be longer than 0"
					areAdminTutorialEditInputsValid[0] = false
				}

				// Description validation
				if len(getAdminTutorialDescriptionEdit) > 0 {
					data.DescriptionError = ""
					areAdminTutorialEditInputsValid[1] = true
				} else {
					data.DescriptionError = "Description should be longer than 0"
					areAdminTutorialEditInputsValid[1] = false
				}

				// URL validation
				if len(getAdminTutorialUrlEdit) > 0 {
					data.UrlError = ""
					areAdminTutorialEditInputsValid[2] = true
				} else {
					data.UrlError = "URL should be longer than 0"
					areAdminTutorialEditInputsValid[2] = false
				}

				// Published validation
				if len(getAdminTutorialPublishedEdit) > 0 {
					data.PublishedError = ""
					areAdminTutorialEditInputsValid[3] = true
				} else {
					data.PublishedError = "Add a date"
					areAdminTutorialEditInputsValid[3] = false
				}

				// Update validation
				if len(getAdminTutorialUpdatedEdit) > 0 {
					data.UpdatedError = ""
					areAdminTutorialEditInputsValid[4] = true
				} else {
					data.UpdatedError = "Add a date"
					areAdminTutorialEditInputsValid[4] = false
				}

				// Image validation
				if len(getAdminTutorialImageEdit) > 0 {
					data.ImageError = ""
					areAdminTutorialEditInputsValid[5] = true
				} else {
					data.ImageError = "An image is required"
					areAdminTutorialEditInputsValid[5] = false
				}

				// Content validation
				if len(getAdminTutorialContentEdit) > 0 {
					data.ContentError = ""
					areAdminTutorialEditInputsValid[6] = true
				} else {
					data.ContentError = "Content should be longher than 0"
					areAdminTutorialEditInputsValid[6] = false
				}

				for i := 0; i < len(areAdminTutorialEditInputsValid); i++ {
					isFormSubmittionValid = true
					if !areAdminTutorialEditInputsValid[i] {
						isFormSubmittionValid = false
						break
					}
				}

				// Edit tutorial if all inputs are valid and redirect to all tutorials list
				if isFormSubmittionValid {
					// Get image id for the relationship one-to-many between tutorials and images
					getAdminTutorialImageIdEdit, _ := models.ImageFindByUrlReturnItsId(getAdminTutorialImageEdit)
					editTutorial := models.TutorialNew(
						tutorialId,
						getAdminTutorialTitleEdit,
						getAdminTutorialDescriptionEdit,
						getAdminTutorialUrlEdit,
						getAdminTutorialPublishedEdit,
						getAdminTutorialUpdatedEdit,
						getAdminTutorialImageIdEdit,
						getAdminTutorialContentEdit,
					)
					models.TutorialEdit(editTutorial)
					http.Redirect(w, r, "/admin/tutorials", http.StatusSeeOther)
				}
			}
			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}

func AdminTutorialDelete() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-tutorial-delete.html"))
	http.HandleFunc("/admin/tutorial-delete/", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["user-admin-authentication"] == true {
			idPath := strings.TrimPrefix(r.URL.Path, "/admin/tutorial-delete/")
			idPath = util.FormSanitizeStringInput(idPath)

			tutorialId, err := strconv.Atoi(idPath)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}

			getTutorialDelete, err := models.TutorialFindById(tutorialId)
			if err != nil {
				fmt.Println("Error to find the tutorial:", err)
			}

			data := tutorialData{
				PageTitle: "Admin Delete Tutorial",
				Tutorials: getTutorialDelete,
			}

			/**
			* Check if the form for deleting tutorial
			* has been submitted
			* and
			* delete the selected tutorial
			 */
			isFormSubmittionValid := false

			// Get the value from the form
			getAdminTutorialDeleteSubmit := r.FormValue("admin-tutorial-delete")

			// Sanitize form input
			getAdminTutorialDeleteSubmit = util.FormSanitizeStringInput(getAdminTutorialDeleteSubmit)

			// Check if the form has been submitted
			if getAdminTutorialDeleteSubmit == "Delete this tutorial" {
				isFormSubmittionValid = true
			}

			if isFormSubmittionValid {
				models.TutorialDelete(tutorialId)
				http.Redirect(w, r, "/admin/tutorials", http.StatusSeeOther)
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
