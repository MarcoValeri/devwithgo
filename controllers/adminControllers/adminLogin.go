package admincontrollers

import (
	"devwithgo/models"
	"fmt"
	"html/template"
	"net/http"
	"net/mail"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

type LoginValidation struct {
	EmailValidation    string
	PasswordValidation string
}

func AdminLogin() {
	tmpl := template.Must(template.ParseFiles("./views/admin/login.html"))
	http.HandleFunc("/admin/login", func(w http.ResponseWriter, r *http.Request) {

		setLoginValidation := LoginValidation{
			EmailValidation:    "",
			PasswordValidation: "",
		}

		// FORM validation
		getEmail := r.FormValue("email")
		getPassword := r.FormValue("password")
		getLogin := r.FormValue("login")

		if len(getLogin) > 0 {
			// Email: check user input
			// Space at the left and right position of the string
			getEmail = strings.TrimSpace(getEmail)

			// Avoid HTML injection
			sanitizeHtml := bluemonday.StrictPolicy()
			getEmail = sanitizeHtml.Sanitize(getEmail)

			// Check if email format is right
			_, err := mail.ParseAddress(getEmail)
			if err != nil {
				setLoginValidation.EmailValidation = "Error: email format is not valid"
				fmt.Println(setLoginValidation.EmailValidation)
			}

			// Check email length, no less that 5 charactes, no longer than 40 characters
			if len(getEmail) < 5 || len(getEmail) > 40 {
				setLoginValidation.EmailValidation = "Error: email length must be greater than 5 characters and no greater than 20 characters"
				fmt.Println(len(getEmail))
			}

			// Password: check user input
			// Remove space to the left and right of the string
			getPassword = strings.TrimSpace(getPassword)

			// Avodi HTML injection
			getPassword = sanitizeHtml.Sanitize(getPassword)

			// Check password length, minimum 8 charactes, max 20
			if len(getPassword) < 8 || len(getPassword) > 20 {
				setLoginValidation.PasswordValidation = "Error: password must be minimum 8 characters but no longher than 20 characters"
			}

			// Form validation
			if models.IsAnUserAdmin(getEmail, getPassword) {
				http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			} else {
				setLoginValidation.EmailValidation = "Error: email and password are not valid"
				setLoginValidation.PasswordValidation = "Error: email and password are not valid"
			}
		}

		tmpl.Execute(w, setLoginValidation)
	})
}
