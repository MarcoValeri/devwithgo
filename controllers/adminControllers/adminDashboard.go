package admincontrollers

import (
	"fmt"
	"html/template"
	"net/http"
)

func AdminDashboard() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-dashboard.html"))
	http.HandleFunc("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		if session.Values["user-admin-authentication"] == true {
			data := "Admin Dashboard"
			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
