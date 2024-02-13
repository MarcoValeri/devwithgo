package admincontrollers

import (
	"fmt"
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

		// Get value form the form
		getAdminUserEmail := r.FormValue("user-add-email")
		getAdminUserPassword := r.FormValue("user-add-password")
		getAdminUserPasswordRepeat := r.FormValue("user-add-password-repeat")
		getAdminUserSubmit := r.FormValue("user-add-new-submit")

		fmt.Println(getAdminUserEmail)
		fmt.Println(getAdminUserPassword)
		fmt.Println(getAdminUserPasswordRepeat)
		fmt.Println(getAdminUserSubmit)

		// TODO: form validation

		data := "Admin Add User"
		tmpl.Execute(w, data)
	})
}
