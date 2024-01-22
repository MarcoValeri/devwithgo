package admincontrollers

import (
	"fmt"
	"html/template"
	"net/http"
)

func AdminLogin() {
	tmpl := template.Must(template.ParseFiles("./views/admin/login.html"))
	http.HandleFunc("/admin/login", func(w http.ResponseWriter, r *http.Request) {

		// FORM validation
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Println("Email:", email)
		fmt.Println("Password:", password)

		tmpl.Execute(w, nil)
	})
}
