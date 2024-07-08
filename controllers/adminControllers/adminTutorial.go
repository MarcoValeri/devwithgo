package admincontrollers

import (
	"devwithgo/models"
	"fmt"
	"html/template"
	"net/http"
)

type tutorialData struct {
	PageTitle string
	Tutorials []models.Tutorial
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
				PageTitle: "Tutorials",
				Tutorials: tutorialsData,
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}

func AdminTutorialAdd() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-tutorial-add.html"))
	http.HandleFunc("admin/tutorial-add", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {
			data := tutorialData{
				PageTitle: "Tutorial Add",
			}

			// Flag validation
			// var areAdminTutorialInputsValid [7]bool
			// isFormSubmittionValid := false

			// Get values from the form
			// getAdminTutorialTitle := r.FormValue("tutorial-title")
			// getAdminTutorialDescription := r.FormValue("tutorial-description")
			// getAdminTutorialUrl := r.FormValue("tutorial-url")
			// getAdminTutorialPublished := r.FormValue("tutorial-published")
			// getAdminTutorialUpdated := r.FormValue("tutorial-updated")
			getAdminTutorialImage := r.FormValue("tutorial-image")
			getAdminTutorialContent := r.FormValue("tutorial-content")
			getAdminTutorialAdd := r.FormValue("tutorial-add")

			// Test
			fmt.Println("Image:", getAdminTutorialImage)
			fmt.Println("Content:", getAdminTutorialContent)
			fmt.Println("Add:", getAdminTutorialAdd)

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
