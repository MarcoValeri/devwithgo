package models

import "devwithgo/database"

type userAdmin struct {
	id       int
	email    string
	password string
}

func UserAdminNew(getUserAdminId int, getUserAdminEmail, getUserAdminPassowrd string) userAdmin {
	setNewUserAdmin := userAdmin{
		id:       getUserAdminId,
		email:    getUserAdminEmail,
		password: getUserAdminPassowrd,
	}
	return setNewUserAdmin
}

func UserAdminAddNewToDB(getNewUserAdmin userAdmin) {
	db := database.DatabaseConnection("local")
	query, err := db.Query("INSERT INTO users (email, password) VALUES (?, ?)", getNewUserAdmin.email, getNewUserAdmin.password)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
}

func IsAnUserAdmin(getEmail, getPassword string) bool {
	if getEmail == "info@marcovaleri.net" && getPassword == "1234" {
		return true
	}
	return false
}
