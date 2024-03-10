package admincontrollers

import (
	"fmt"
	"html/template"
	"net/http"
)

type guideData struct {
	PageTitle string
	Guides    string
}

func AdminGuides() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-guides.html"))
	http.HandleFunc("/admin/guides", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {
			// TODO: get and show all guides
			data := guideData{
				PageTitle: "Guides",
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
			// var areAdminGuideInputsValid [6]bool
			// isFormSubmittionValid := false

			// Get value from the form
			getAdminGuideTitle := r.FormValue("guide-title")
			getAdminGuideDescription := r.FormValue("guide-description")
			getAdminGuideUrl := r.FormValue("guide-url")
			getAdminGuidePublished := r.FormValue("guide-published")
			getAdminGuideUpdated := r.FormValue("guide-updated")
			getAdminGuideContent := r.FormValue("guide-content")
			getAdminGuideAdd := r.FormValue("guide-add")

			fmt.Println(getAdminGuideTitle)
			fmt.Println(getAdminGuideDescription)
			fmt.Println(getAdminGuideUrl)
			fmt.Println(getAdminGuidePublished)
			fmt.Println(getAdminGuideUpdated)
			fmt.Println(getAdminGuideContent)
			fmt.Println(getAdminGuideAdd)

			// TODO: validate form and add the guide to the db

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}
