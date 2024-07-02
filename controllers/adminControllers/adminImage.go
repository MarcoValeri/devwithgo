package admincontrollers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func AdminUploadImage() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-image-add.html"))
	http.HandleFunc("/admin/image-add", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {
			// TODO: store image data into the db and upload image into its folder

			// Form inputs
			imageAddNew := r.FormValue("image-add-new")
			imageFile, header, errImageFile := r.FormFile("image-file")

			if len(imageAddNew) > 0 {
				if errImageFile != nil {
					fmt.Println("Error retrieving the image file:", errImageFile)
					return
				}
				// defer imageFile.Close()

				imagePath := filepath.Join("public", "images", header.Filename)
				absImagePath, errImagePath := filepath.Abs(imagePath)
				if errImagePath != nil {
					fmt.Println("Error determing image path:", errImagePath)
					return
				}

				dst, erDst := os.Create(absImagePath)
				if erDst != nil {
					fmt.Println("Error creating image file:", erDst)
				}

				_, errCopy := io.Copy(dst, imageFile)
				if errCopy != nil {
					fmt.Println("Error saving image file:", errCopy)
					return
				}

				defer dst.Close()
				defer imageFile.Close()

				http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			}

			tmpl.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}
