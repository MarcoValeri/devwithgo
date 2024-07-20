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
	"strconv"
	"strings"
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
	Images                []models.Image
}

func AdminImages() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-images.html"))
	http.HandleFunc("/admin/images", func(w http.ResponseWriter, r *http.Request) {
		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {
			imagesData, err := models.ImageShowImages()
			if err != nil {
				fmt.Println("Error getting imagesData:", err)
			}

			data := imageData{
				PageTitle: "Admin Images",
				Images:    imagesData,
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}

func AdminImageAdd() {
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
							getImageCredit,
							getImageUrl,
							getImagePublished,
							getImageUpdated,
						)
						models.ImageAddNewToDB(createNewImage)
						http.Redirect(w, r, "/admin/images", http.StatusSeeOther)
					}

				}
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}

func AdminImageEdit() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-image-edit.html"))
	http.HandleFunc("/admin/image-edit/", func(w http.ResponseWriter, r *http.Request) {
		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["user-admin-authentication"] == true {
			idPath := strings.TrimPrefix(r.URL.Path, "/admin/image-edit/")
			idPath = util.FormSanitizeStringInput(idPath)

			imageId, err := strconv.Atoi(idPath)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}

			getImageEdit, err := models.ImageFindIt(imageId)
			if err != nil {
				fmt.Println("Error to find image by id:", err)
			}

			data := imageData{
				PageTitle: "Admin Image Edit",
				Images:    getImageEdit,
			}

			/**
			* Check if the form for editing the image has been submitted
			* and
			* validate the inputs
			 */
			var areAdminImageEditInputsValid [6]bool
			isFormSubmittionValid := false

			// Get the values from the form
			getAdminImageTitleEdit := r.FormValue("image-title-edit")
			getAdminImageDescriptionEdit := r.FormValue("image-description-edit")
			getAdminImageCreditEdit := r.FormValue("image-credit-edit")
			getAdminImageUrlEdit := r.FormValue("image-url-edit")
			getAdminImagePublishedEdit := r.FormValue("image-published-edit")
			getAdminImageUpdatedEdit := r.FormValue("image-updated-edit")
			getAdminImageSubmitEdit := r.FormValue("image-edit")

			// Sanitize form inputs
			getAdminImageTitleEdit = util.FormSanitizeStringInput(getAdminImageTitleEdit)
			getAdminImageDescriptionEdit = util.FormSanitizeStringInput(getAdminImageDescriptionEdit)
			getAdminImageCreditEdit = util.FormSanitizeStringInput(getAdminImageCreditEdit)
			getAdminImageUrlEdit = util.FormSanitizeStringInput(getAdminImageUrlEdit)
			getAdminImagePublishedEdit = util.FormSanitizeStringInput(getAdminImagePublishedEdit)
			getAdminImageUpdatedEdit = util.FormSanitizeStringInput(getAdminImageUpdatedEdit)
			getAdminImageSubmitEdit = util.FormSanitizeStringInput(getAdminImageSubmitEdit)

			// Check if the form has been submitted
			if getAdminImageSubmitEdit == "Edit this image" {
				// Title validation
				if len(getAdminImageTitleEdit) > 0 {
					data.ImageTitleError = ""
					areAdminImageEditInputsValid[0] = true
				} else {
					data.ImageTitleError = "Title should be longer than 0 characters"
					areAdminImageEditInputsValid[0] = false
				}

				// Description validation
				if len(getAdminImageDescriptionEdit) > 0 {
					data.ImageDescriptionError = ""
					areAdminImageEditInputsValid[1] = true
				} else {
					data.ImageDescriptionError = "Description should be longer than 0 characters"
					areAdminImageEditInputsValid[1] = false
				}

				// Credit validation
				if len(getAdminImageCreditEdit) > 0 {
					data.ImageCreditError = ""
					areAdminImageEditInputsValid[2] = true
				} else {
					data.ImageCreditError = "Credit should be longer than 0 characters"
					areAdminImageEditInputsValid[2] = false
				}

				// Url validation
				if len(getAdminImageUrlEdit) > 0 {
					data.ImageUrlError = ""
					areAdminImageEditInputsValid[3] = true
				} else {
					data.ImageUrlError = "URL should be longer than 0 characters"
					areAdminImageEditInputsValid[3] = false
				}

				// Published validation
				if len(getAdminImagePublishedEdit) > 0 {
					data.ImagePublishedError = ""
					areAdminImageEditInputsValid[4] = true
				} else {
					data.ImagePublishedError = "Add a date"
					areAdminImageEditInputsValid[4] = false
				}

				// Uploaded validation
				if len(getAdminImageUpdatedEdit) > 0 {
					data.ImageUpdatedError = ""
					areAdminImageEditInputsValid[5] = true
				} else {
					data.ImageUpdatedError = "Add a date"
					areAdminImageEditInputsValid[5] = false
				}

				for i := 0; i < len(areAdminImageEditInputsValid); i++ {
					isFormSubmittionValid = true
					if !areAdminImageEditInputsValid[i] {
						isFormSubmittionValid = false
						break
					}
				}

				// Edit image if all inputs are valid and redirect to all images list
				if isFormSubmittionValid {
					editImage := models.ImageNew(imageId, getAdminImageTitleEdit, getAdminImageDescriptionEdit, getAdminImageCreditEdit, getAdminImageUrlEdit, getAdminImagePublishedEdit, getAdminImageUpdatedEdit)
					models.ImageEdit(editImage)
					http.Redirect(w, r, "/admin/images", http.StatusSeeOther)
				}
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}

func AdminImageDelete() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-image-delete.html"))
	http.HandleFunc("/admin/image-delete/", func(w http.ResponseWriter, r *http.Request) {
		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {
			idPath := strings.TrimPrefix(r.URL.Path, "/admin/image-delete/")
			idPath = util.FormSanitizeStringInput(idPath)

			imageId, err := strconv.Atoi(idPath)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}

			getImageDelete, err := models.ImageFindIt(imageId)
			if err != nil {
				fmt.Println("Error to find image:", err)
			}

			data := imageData{
				PageTitle: "Admin Delete Image",
				Images:    getImageDelete,
			}

			/**
			* Check if the form for deleting image has
			* been submitted
			* and
			* delete the selected image
			 */
			isFormSubmittionValid := false

			// Get the value from the form
			getAdminImageDeleteUrl := r.FormValue("admin-image-url-delete")
			getAdminImageDeleteSubmit := r.FormValue("admin-image-delete")

			// Sanitize form input
			getAdminImageDeleteUrl = util.FormSanitizeStringInput(getAdminImageDeleteUrl)
			getAdminImageDeleteSubmit = util.FormSanitizeStringInput(getAdminImageDeleteSubmit)

			// Check if the form has been submitted
			if getAdminImageDeleteSubmit == "Delete this image" && len(getAdminImageDeleteUrl) > 0 {

				// TODO: delete image from the images folder
				imagePath := filepath.Join("public/images", getAdminImageDeleteUrl)

				if _, err := os.Stat(imagePath); os.IsNotExist(err) {
					fmt.Println("Image not found:", err)
					isFormSubmittionValid = false
				} else {
					isFormSubmittionValid = true
				}

				err := os.Remove(imagePath)
				if err != nil {
					fmt.Println("Error deleteing image:", err)
					isFormSubmittionValid = false
				} else {
					isFormSubmittionValid = true
				}

			}

			if isFormSubmittionValid {
				models.ImageDelete(imageId)
				http.Redirect(w, r, "/admin/images", http.StatusSeeOther)
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
