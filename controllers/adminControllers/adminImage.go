package admincontrollers

import (
	"devwithgo/models"
	"devwithgo/util"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type imageData struct {
	PageTitle             string
	ImageTitleError       string
	ImageUrlError         string
	ImageDescriptionError string
	ImageCreditError      string
	ImagePublishedError   string
	ImageUpdatedError     string
	ImageFileError        string
}

func AdminUploadImage() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-image-add.html"))
	http.HandleFunc("/admin/image-add", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {

			data := imageData{
				PageTitle: "Admin Add Image",
			}

			// Flag validation
			var areAdminImageInputsValid [7]bool
			isFormSubmittionValid := false

			// Get values from the form
			getImageTitle := r.FormValue("image-title")
			getImageUrl := r.FormValue("image-url")
			getImageDescription := r.FormValue("image-description")
			getImageCredit := r.FormValue("image-credit")
			getImagePublished := r.FormValue("image-published")
			getImageUpdated := r.FormValue("image-updated")
			getImageAddNew := r.FormValue("image-add-new")
			getImageFile, header, errImageFile := r.FormFile("image-file")

			// Sanitize form inputs
			getImageTitle = util.FormSanitizeStringInput(getImageTitle)
			getImageUrl = util.FormSanitizeStringInput(getImageUrl)
			getImageDescription = util.FormSanitizeStringInput(getImageDescription)
			getImageCredit = util.FormSanitizeStringInput(getImageCredit)
			getImagePublished = util.FormSanitizeStringInput(getImagePublished)
			getImageUpdated = util.FormSanitizeStringInput(getImageUpdated)
			getImageAddNew = util.FormSanitizeStringInput(getImageAddNew)

			if getImageAddNew == "Add new image" {
				// Image Title validation
				if len(getImageTitle) > 0 {
					data.ImageTitleError = ""
					areAdminImageInputsValid[0] = true
				} else {
					data.ImageTitleError = "Title should be longer than 0 characters"
					areAdminImageInputsValid[0] = false
				}

				// Image Url validation
				if len(getImageUrl) > 0 {
					data.ImageUrlError = ""
					areAdminImageInputsValid[1] = true
				} else {
					data.ImageUrlError = "URL should be longer than 0 characters"
					areAdminImageInputsValid[1] = false
				}

				// Image description validation
				if len(getImageDescription) > 0 {
					data.ImageDescriptionError = ""
					areAdminImageInputsValid[2] = true
				} else {
					data.ImageDescriptionError = "Description should be longer than 0 characters"
					areAdminImageInputsValid[2] = false
				}

				// Image credit validation
				if len(getImageCredit) > 0 {
					data.ImageCreditError = ""
					areAdminImageInputsValid[3] = true
				} else {
					data.ImageDescriptionError = "Credit should be longer than 0 characters"
					areAdminImageInputsValid[3] = false
				}

				// Image Published validation
				if len(getImagePublished) > 0 {
					data.ImagePublishedError = ""
					areAdminImageInputsValid[4] = true
				} else {
					data.ImagePublishedError = "Inser a valid date"
					areAdminImageInputsValid[4] = false
				}

				// Image Updated validation
				if len(getImageUpdated) > 0 {
					data.ImagePublishedError = ""
					areAdminImageInputsValid[5] = true
				} else {
					data.ImagePublishedError = "Inser a valid date"
					areAdminImageInputsValid[5] = false
				}

				// Image file validation
				if util.FormIsValidImage(getImageFile, header.Filename) {
					data.ImageFileError = ""
					areAdminImageInputsValid[6] = true
				} else {
					data.ImageFileError = "Image file is not valid"
					areAdminImageInputsValid[6] = false
				}

				for i := 0; i < len(areAdminImageInputsValid); i++ {
					isFormSubmittionValid = true
					if !areAdminImageInputsValid[i] {
						isFormSubmittionValid = false
						break
					}
				}

				// Store image and save its data to the db
				if isFormSubmittionValid {

					// Flag validation for uploading image
					var isImageUploadedCorrectly [4]bool
					isImageUploaded := false

					if errImageFile != nil {
						fmt.Println("Error retrieving the image file:", errImageFile)
						data.ImageFileError = "Image file is not valid"
						isImageUploadedCorrectly[0] = false
					} else {
						data.ImageFileError = ""
						isImageUploadedCorrectly[0] = true
					}

					imagePath := filepath.Join("public", "images", header.Filename)
					absImagePath, errImagePath := filepath.Abs(imagePath)
					if errImagePath != nil {
						fmt.Println("Error determing image path:", errImagePath)
						data.ImageFileError = "Image file is not valid"
						isImageUploadedCorrectly[1] = false
					} else {
						data.ImageFileError = ""
						isImageUploadedCorrectly[1] = true
					}

					dst, erDst := os.Create(absImagePath)
					if erDst != nil {
						fmt.Println("Error creating image file:", erDst)
						data.ImageFileError = "Image file is not valid"
						isImageUploadedCorrectly[2] = false
					} else {
						data.ImageFileError = ""
						isImageUploadedCorrectly[2] = true
					}

					_, errCopy := io.Copy(dst, getImageFile)
					if errCopy != nil {
						fmt.Println("Error saving image file:", errCopy)
						data.ImageFileError = "Image file is not valid"
						isImageUploadedCorrectly[3] = false
					} else {
						data.ImageFileError = ""
						isImageUploadedCorrectly[3] = true
					}

					defer dst.Close()
					defer getImageFile.Close()

					for i := 0; i < len(isImageUploadedCorrectly); i++ {
						isImageUploaded = true
						if !isImageUploadedCorrectly[i] {
							isImageUploaded = false
							break
						}
					}

					if isImageUploaded {
						createNewImage := models.ImageNew(
							1,
							getImageTitle,
							getImageDescription,
							getImageUrl,
							getImagePublished,
							getImageUpdated,
						)
						models.ImageAddNewToDB(createNewImage)
						http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
					}

				}
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}
