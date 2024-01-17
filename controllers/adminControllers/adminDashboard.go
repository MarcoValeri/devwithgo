package admincontrollers

import (
	"html/template"
	"net/http"
)

func AdminDashboard() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/dashboard.html"))
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		data := "Admin Dashboard"
		tmpl.Execute(w, data)
	})
}
