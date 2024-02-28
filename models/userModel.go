package models

import (
	"devwithgo/database"
	"fmt"
	"strconv"
)

type UserAdmin struct {
	Id       int
	Email    string
	Password string
}

func UserAdminNew(getUserAdminId int, getUserAdminEmail, getUserAdminPassowrd string) UserAdmin {
	setNewUserAdmin := UserAdmin{
		Id:       getUserAdminId,
		Email:    getUserAdminEmail,
		Password: getUserAdminPassowrd,
	}
	return setNewUserAdmin
}

func UserAdminAddNewToDB(getNewUserAdmin UserAdmin) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query("INSERT INTO users (email, password) VALUES (?, ?)", getNewUserAdmin.Email, getNewUserAdmin.Password)
	if err != nil {
		return fmt.Errorf("error adding user: %w", err)
	}
	defer query.Close()

	return nil
}

func UserEdminEdit(getEditedUserAdmin UserAdmin) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query("UPDATE users SET email = ?, password = ? WHERE id = ?", getEditedUserAdmin.Email, getEditedUserAdmin.Password, getEditedUserAdmin.Id)
	if err != nil {
		fmt.Println("Error on editing user query")
		return err
	}
	defer query.Close()

	return nil
}

func UserAdminFindIt(getUserAdminId int) ([][]string, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE id=?", getUserAdminId)
	if err != nil {
		fmt.Println("Error on the user query")
		return nil, err
	}
	defer rows.Close()

	var getUserData [][]string
	for rows.Next() {
		var userId int
		var userEmail string
		var userPw string
		err = rows.Scan(&userId, &userEmail, &userPw)
		if err != nil {
			return nil, err
		}
		userDetails := []string{strconv.Itoa(userId), userEmail, userPw}
		getUserData = append(getUserData, userDetails)
	}

	return getUserData, nil
}

func IsAnUserAdmin(getEmail, getPassword string) bool {
	if getEmail == "info@marcovaleri.net" && getPassword == "1234" {
		return true
	}
	return false
}

func UserAdminShowUsers() ([]UserAdmin, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// var allUsers [][]string
	// for rows.Next() {
	// 	var userId int
	// 	var userEmail string
	// 	var userPw string
	// 	err = rows.Scan(&userId, &userEmail, &userPw)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	userDatails := []string{strconv.Itoa(userId), userEmail, userPw}
	// 	allUsers = append(allUsers, userDatails)
	// }

	var allUsers []UserAdmin
	for rows.Next() {
		var userId int
		var userEmail string
		var userPw string
		err = rows.Scan(&userId, &userEmail, &userPw)
		if err != nil {
			return nil, err
		}
		// userDatails := []string{strconv.Itoa(userId), userEmail, userPw}
		userDatails := UserAdminNew(userId, userEmail, userPw)
		allUsers = append(allUsers, userDatails)
	}

	fmt.Println(allUsers)

	return allUsers, nil
}
