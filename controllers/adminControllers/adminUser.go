package admincontrollers

import (
	"devwithgo/models"
	"devwithgo/util"
	"html/template"
	"net/http"
)

type dataPage struct {
	PageTitle           string
	EmailError          string
	PasswordError       string
	PasswordRepearError string
	PasswordMatch       string
}

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

		data := dataPage{
			PageTitle: "Dashboard Admin Add User",
		}

		// Flag validation
		var areAdminUserInputsValid [5]bool
		isFormSubmittionValid := false

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
				data.EmailError = ""
				areAdminUserInputsValid[0] = true
				if util.FormEmailLengthInput(getAdminUserEmail) && areAdminUserInputsValid[0] {
					data.EmailError = ""
					areAdminUserInputsValid[0] = true
				} else {
					data.EmailError = "Email length is not valid"
					areAdminUserInputsValid[0] = false
				}
			} else {
				data.EmailError = "Email format is not valid"
				areAdminUserInputsValid[0] = false
			}

			// Password validation
			if util.FormPasswordInput(getAdminUserPassword) {
				data.PasswordError = ""
				areAdminUserInputsValid[1] = true
			} else {
				data.PasswordError = "Password should be between 8 to 20 characters"
				areAdminUserInputsValid[1] = false
			}

			if util.FormPasswordInput(getAdminUserPasswordRepeat) {
				data.PasswordRepearError = ""
				areAdminUserInputsValid[2] = true
			} else {
				data.PasswordRepearError = "Password should be between 8 to 20 characters"
				areAdminUserInputsValid[2] = false
			}

			if getAdminUserPassword == getAdminUserPasswordRepeat {
				data.PasswordMatch = ""
				areAdminUserInputsValid[3] = true
			} else {
				data.PasswordMatch = "Password and repeat password do not match"
				areAdminUserInputsValid[3] = false
			}

			// Submit validation
			if getAdminUserSubmit == "Add new user" {
				areAdminUserInputsValid[4] = true
			} else {
				areAdminUserInputsValid[4] = false
			}

			for i := 0; i < len(areAdminUserInputsValid); i++ {
				isFormSubmittionValid = true
				if !areAdminUserInputsValid[i] {
					isFormSubmittionValid = false
					break
				}
			}

			// Create a new user if all inputs are valid
			if isFormSubmittionValid {
				createNewUserAdmin := models.UserAdminNew(1, getAdminUserEmail, getAdminUserPassword)
				models.UserAdminAddNewToDB(createNewUserAdmin)
				http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			}
		}

		tmpl.Execute(w, data)
	})
}
