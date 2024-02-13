package admincontrollers

import (
	"html/template"
	"net/http"
)

func AdminUsers() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-users.html"))
	http.HandleFunc("/admin/users", func(w http.ResponseWriter, r *http.Request) {
		data := "Admin Users"
		tmpl.Execute(w, data)
	})
}

func AdminUserAdd() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-user-add.html"))
	http.HandleFunc("/admin/user-add", func(w http.ResponseWriter, r *http.Request) {
		data := "Admin Add User"
		tmpl.Execute(w, data)
	})
}
