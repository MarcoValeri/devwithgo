package models

import (
	"devwithgo/database"
	"fmt"
	"strconv"
)

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

func UserAdminAddNewToDB(getNewUserAdmin userAdmin) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query("INSERT INTO users (email, password) VALUES (?, ?)", getNewUserAdmin.email, getNewUserAdmin.password)
	if err != nil {
		return fmt.Errorf("error adding user: %w", err)
	}
	defer query.Close()

	return nil
}

func IsAnUserAdmin(getEmail, getPassword string) bool {
	if getEmail == "info@marcovaleri.net" && getPassword == "1234" {
		return true
	}
	return false
}

func UserAdminShowUsers() ([][]string, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allUsers [][]string
	for rows.Next() {
		var userId int
		var userEmail string
		var userPw string
		err = rows.Scan(&userId, &userEmail, &userPw)
		if err != nil {
			return nil, err
		}
		userDatails := []string{strconv.Itoa(userId), userEmail, userPw}
		allUsers = append(allUsers, userDatails)
	}

	return allUsers, nil
}
