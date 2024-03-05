package models

import (
	"devwithgo/database"
	"devwithgo/util"
	"fmt"
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

	hashThePassword, errHashPassword := util.PasswordHash(getNewUserAdmin.Password)
	if errHashPassword != nil {
		fmt.Println("Error to hash the password:", errHashPassword)
	}

	query, err := db.Query("INSERT INTO users (email, password) VALUES (?, ?)", getNewUserAdmin.Email, hashThePassword)
	if err != nil {
		return fmt.Errorf("error adding user: %w", err)
	}
	defer query.Close()

	return nil
}

func UserAdminEdit(getEditedUserAdmin UserAdmin) error {
	db := database.DatabaseConnection()
	defer db.Close()

	hashThePassword, errHashPassword := util.PasswordHash(getEditedUserAdmin.Password)
	if errHashPassword != nil {
		fmt.Println("Error to hash the password:", errHashPassword)
	}

	query, err := db.Query("UPDATE users SET email = ?, password = ? WHERE id = ?", getEditedUserAdmin.Email, hashThePassword, getEditedUserAdmin.Id)
	if err != nil {
		fmt.Println("Error on editing user query")
		return err
	}
	defer query.Close()

	return nil
}

func UserAdminDelete(getUserAdminId int) error {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("DELETE FROM users WHERE id=?", getUserAdminId)
	if err != nil {
		fmt.Println("Error, not able to delete this user:", err)
		return err
	}
	defer rows.Close()

	return nil
}

func UserAdminFindIt(getUserAdminId int) ([]UserAdmin, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE id=?", getUserAdminId)
	if err != nil {
		fmt.Println("Error on the user query")
		return nil, err
	}
	defer rows.Close()

	var getUserData []UserAdmin
	for rows.Next() {
		var userId int
		var userEmail string
		var userPw string
		err = rows.Scan(&userId, &userEmail, &userPw)
		if err != nil {
			return nil, err
		}
		userDatails := UserAdminNew(userId, userEmail, userPw)
		getUserData = append(getUserData, userDatails)
	}

	return getUserData, nil
}

func UserAdminLogin(getEmail, getPassword string) bool {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE email=?", getEmail)
	if err != nil {
		fmt.Println("Error to user admin logic query:", err)
		return false
	}
	defer rows.Close()

	var getUserAdminEmail string
	var getUserAdminPassword string
	for rows.Next() {
		var userId int
		var userEmail string
		var userPassword string
		err = rows.Scan(&userId, &userEmail, &userPassword)
		if err != nil {
			fmt.Println("Error to user admin logic fetching data:", err)
			return false
		}
		getUserAdminEmail = userEmail
		getUserAdminPassword = userPassword
	}

	if len(getUserAdminEmail) > 0 && len(getUserAdminPassword) > 0 {
		userAdminPasswordMath := util.PasswordHashChecker(getPassword, getUserAdminPassword)
		if userAdminPasswordMath {
			return true
		}
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

	var allUsers []UserAdmin
	for rows.Next() {
		var userId int
		var userEmail string
		var userPw string
		err = rows.Scan(&userId, &userEmail, &userPw)
		if err != nil {
			return nil, err
		}
		userDatails := UserAdminNew(userId, userEmail, userPw)
		allUsers = append(allUsers, userDatails)
	}

	return allUsers, nil
}
