package admincontrollers

import (
	"devwithgo/models"
	"devwithgo/util"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

type LoginValidation struct {
	EmailValidation    string
	PasswordValidation string
}

// Initialize the session
var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY_LOGIN")))

func AdminLogin() {
	tmpl := template.Must(template.ParseFiles("./views/admin/admin-login.html"))
	http.HandleFunc("/admin/login", func(w http.ResponseWriter, r *http.Request) {

		setLoginValidation := LoginValidation{
			EmailValidation:    "",
			PasswordValidation: "",
		}

		// Session authentication
		session, errSession := store.Get(r, "session-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}
		session.Values["user-admin-authentication"] = true
		session.Save(r, w)
		fmt.Println(session.Values["user-admin-authentication"])
		fmt.Println(os.Getenv("SESSION_KEY_LOGIN"))

		// FORM validation
		getEmail := r.FormValue("email")
		getPassword := r.FormValue("password")
		getLogin := r.FormValue("login")

		if len(getLogin) > 0 {
			getEmail = util.FormSanitizeStringInput(getEmail)
			getPassword = util.FormSanitizeStringInput(getPassword)

			// Check if email format is right
			if !util.FormEmailInput(getEmail) {
				setLoginValidation.EmailValidation = "Error: email format is not valid"
			}

			// Check email length, no less that 5 charactes, no longer than 40 characters
			if !util.FormEmailLengthInput(getEmail) {
				setLoginValidation.EmailValidation = "Error: email length must be greater than 5 characters and no greater than 20 characters"
			}

			// Check password length, minimum 8 charactes, max 20
			if !util.FormPasswordInput(getPassword) {
				setLoginValidation.PasswordValidation = "Error: password must be minimum 8 characters but no longher than 20 characters"
			}

			// Form validation
			if models.UserAdminLogin(getEmail, getPassword) {
				// TODO: create login session

				http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			} else {
				setLoginValidation.EmailValidation = "Error: email and password are not valid"
				setLoginValidation.PasswordValidation = "Error: email and password are not valid"
			}
		}

		tmpl.Execute(w, setLoginValidation)
	})
}
