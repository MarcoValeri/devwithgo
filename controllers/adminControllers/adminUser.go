package admincontrollers

import (
	"devwithgo/models"
	"devwithgo/util"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type dataPage struct {
	PageTitle           string
	EmailError          string
	PasswordError       string
	PasswordRepearError string
	PasswordMatch       string
	UsersData           [][]string
	UserAdminData       []models.UserAdmin
}

func AdminUsers() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-users.html"))
	http.HandleFunc("/admin/users", func(w http.ResponseWriter, r *http.Request) {

		usersData, err := models.UserAdminShowUsers()
		if err != nil {
			fmt.Println("error getting usetsData: %w", err)
		}
		fmt.Println(usersData)

		data := dataPage{
			PageTitle:     "Admin Users",
			UserAdminData: usersData,
		}

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

		// Get value from the form
		getAdminUserEmail := r.FormValue("user-add-email")
		getAdminUserPassword := r.FormValue("user-add-password")
		getAdminUserPasswordRepeat := r.FormValue("user-add-password-repeat")
		getAdminUserSubmit := r.FormValue("user-add-new-submit")

		// Sanitize form inputs
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
				http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
			}
		}

		tmpl.Execute(w, data)
	})
}

func AdminUserEdit() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-user-edit.html"))
	http.HandleFunc("/admin/user-edit/", func(w http.ResponseWriter, r *http.Request) {

		idPath := strings.TrimPrefix(r.URL.Path, "/admin/user-edit/")
		idPath = util.FormSanitizeStringInput(idPath)

		userId, err := strconv.Atoi(idPath)
		if err != nil {
			fmt.Println("Error convertin string to integer:", err)
			return
		}

		getUserEdit, err := models.UserAdminFindIt(userId)
		if err != nil {
			fmt.Println("Error to find user")
		}

		// Create data for the page
		data := dataPage{
			PageTitle: "Admin User Edit",
			UsersData: getUserEdit,
		}

		/**
		* Check if the form for editing the user has been submitted
		* and
		* validate the inputs
		 */
		var areAdminUserEditInputsValid [5]bool
		isFormSubmittionValid := false

		// Get value from the form
		getAdminUserEmailEdit := r.FormValue("user-admin-email-edit")
		getAdminUserPasswordEdit := r.FormValue("user-admin-password-edit")
		getAdminUserPassordRepeatEdit := r.FormValue("user-admin-password-repeat-edit")
		getAdminUserSubmitEdit := r.FormValue("user-admin-edit-submit")

		// Sanitize form inputs
		getAdminUserEmailEdit = util.FormSanitizeStringInput(getAdminUserEmailEdit)
		getAdminUserPasswordEdit = util.FormSanitizeStringInput(getAdminUserPasswordEdit)
		getAdminUserPassordRepeatEdit = util.FormSanitizeStringInput(getAdminUserPassordRepeatEdit)
		getAdminUserSubmitEdit = util.FormSanitizeStringInput(getAdminUserSubmitEdit)

		// Check if the form has been submitted
		if getAdminUserSubmitEdit == "Edit this user" {
			// Email validation
			if util.FormEmailInput(getAdminUserEmailEdit) {
				data.EmailError = ""
				areAdminUserEditInputsValid[0] = true
				if util.FormEmailLengthInput(getAdminUserEmailEdit) && areAdminUserEditInputsValid[0] {
					data.EmailError = ""
					areAdminUserEditInputsValid[0] = true
				} else {
					data.EmailError = "Email length is not valid"
					areAdminUserEditInputsValid[0] = false
				}
			} else {
				data.EmailError = "Email format is not valid"
				areAdminUserEditInputsValid[0] = false
			}
		}

		// Password validation
		if util.FormPasswordInput(getAdminUserPasswordEdit) {
			data.PasswordError = ""
			areAdminUserEditInputsValid[1] = true
		} else {
			data.PasswordError = "Password should be between 8 to 20 characters"
			areAdminUserEditInputsValid[1] = false
		}

		if util.FormPasswordInput(getAdminUserPassordRepeatEdit) {
			data.PasswordRepearError = ""
			areAdminUserEditInputsValid[2] = true
		} else {
			data.PasswordRepearError = "Password should be between 8 to 20 characters"
			areAdminUserEditInputsValid[2] = false
		}

		if getAdminUserPasswordEdit == getAdminUserPassordRepeatEdit {
			data.PasswordMatch = ""
			areAdminUserEditInputsValid[3] = true
		} else {
			data.PasswordMatch = "Password and repeat password do not match"
			areAdminUserEditInputsValid[3] = false
		}

		// Submit validation
		if getAdminUserSubmitEdit == "Edit this user" {
			areAdminUserEditInputsValid[4] = true
		} else {
			areAdminUserEditInputsValid[4] = false
		}

		for i := 0; i < len(areAdminUserEditInputsValid); i++ {
			isFormSubmittionValid = true
			if !areAdminUserEditInputsValid[i] {
				isFormSubmittionValid = false
				break
			}
		}

		// Edit user if all inputs are valid and redirect to all user list
		if isFormSubmittionValid {
			editUserAdmin := models.UserAdminNew(userId, getAdminUserEmailEdit, getAdminUserPasswordEdit)
			models.UserEdminEdit(editUserAdmin)
			http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
		}

		tmpl.Execute(w, data)
	})
}
