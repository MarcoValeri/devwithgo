package admincontrollers

import (
	"devwithgo/util"
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

		// Data struct for the page render
		type dataPage struct {
			PageTitle           string
			EmailError          string
			PasswordError       string
			PasswordRepearError string
		}

		data := dataPage{
			PageTitle: "Dashboard Admin Add User",
		}

		// Flag validation
		var areAdminUserInputsValid [5]bool

		// Get value form the form
		getAdminUserEmail := r.FormValue("user-add-email")
		getAdminUserPassword := r.FormValue("user-add-password")
		getAdminUserPasswordRepeat := r.FormValue("user-add-password-repeat")
		getAdminUserSubmit := r.FormValue("user-add-new-submit")

		// Sanitize form input
		getAdminUserEmail = util.FormSanitizeStringInput(getAdminUserEmail)
		getAdminUserPassword = util.FormSanitizeStringInput(getAdminUserPassword)
		getAdminUserPasswordRepeat = util.FormSanitizeStringInput(getAdminUserPasswordRepeat)
		getAdminUserSubmit = util.FormSanitizeStringInput(getAdminUserSubmit)

		// Check if the form has been submittet
		if getAdminUserSubmit == "Add new user" {
			// Email validation
			if util.FormEmailInput(getAdminUserEmail) {
				areAdminUserInputsValid[0] = true
				if util.FormEmailLengthInput(getAdminUserEmail) && areAdminUserInputsValid[0] {
					areAdminUserInputsValid[0] = true
				} else {
					areAdminUserInputsValid[0] = false
				}
			} else {
				areAdminUserInputsValid[0] = false
			}

			// Password validation
			if util.FormPasswordInput(getAdminUserPassword) {
				areAdminUserInputsValid[1] = true
			} else {
				areAdminUserInputsValid[1] = false
			}

			if util.FormPasswordInput(getAdminUserPasswordRepeat) {
				areAdminUserInputsValid[2] = true
			} else {
				areAdminUserInputsValid[2] = false
			}

			if getAdminUserPassword == getAdminUserPasswordRepeat {
				areAdminUserInputsValid[3] = true
			} else {
				areAdminUserInputsValid[3] = false
			}

			// Submit validation
			if getAdminUserSubmit == "Add new user" {
				areAdminUserInputsValid[4] = true
			} else {
				areAdminUserInputsValid[4] = false
			}

			// TODO: add input error to data

			// TEST VALIDATION
			fmt.Println(getAdminUserEmail)
			fmt.Println(getAdminUserPassword)
			fmt.Println(getAdminUserPasswordRepeat)
			fmt.Println(getAdminUserSubmit)
			fmt.Println(areAdminUserInputsValid, "\n")
		}

		tmpl.Execute(w, data)
	})
}
