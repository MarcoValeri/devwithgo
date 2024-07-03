package models

import (
	"devwithgo/database"
	"fmt"
)

type Image struct {
	Id          int
	Title       string
	Description string
	Url         string
	Published   string
	Updated     string
}

func ImageNew(getImageId int, getImageTitle, getImageDescription, getImageUrl, getImagePublished, getImageUpdated string) Image {
	setNewImage := Image{
		Id:          getImageId,
		Title:       getImageTitle,
		Description: getImageDescription,
		Url:         getImageUrl,
		Published:   getImagePublished,
		Updated:     getImageUpdated,
	}
	return setNewImage
}

func ImageAddNewToDB(getImage Image) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query(
		"INSERT INTO images (title, description, url, published, updated) VALUES (?, ?, ?, ?, ?)",
		getImage.Title, getImage.Description, getImage.Url, getImage.Published, getImage.Updated,
	)
	if err != nil {
		fmt.Println("Error adding new image data to the db:", err)
		return err
	}
	defer query.Close()

	return nil
}
