package models

import (
	"devwithgo/database"
	"fmt"
)

type Image struct {
	Id          int
	Title       string
	Description string
	Credit      string
	Url         string
	Published   string
	Updated     string
}

func ImageNew(getImageId int, getImageTitle, getImageDescription, getCredit, getImageUrl, getImagePublished, getImageUpdated string) Image {
	setNewImage := Image{
		Id:          getImageId,
		Title:       getImageTitle,
		Description: getImageDescription,
		Credit:      getCredit,
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
		"INSERT INTO images (title, description, credit, url, published, updated) VALUES (?, ?, ?, ?, ?, ?)",
		getImage.Title, getImage.Description, getImage.Credit, getImage.Url, getImage.Published, getImage.Updated,
	)
	if err != nil {
		fmt.Println("Error adding new image data to the db:", err)
		return err
	}
	defer query.Close()

	return nil
}

func ImageShowImages() ([]Image, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM images")
	if err != nil {
		fmt.Println("Error retriving images data from the db:", err)
		return nil, err
	}
	defer rows.Close()

	var allImages []Image
	for rows.Next() {
		var imageId int
		var imageTitle string
		var imageDescription string
		var imageCredit string
		var imageUrl string
		var imagePublished string
		var imageUpdated string
		err = rows.Scan(&imageId, &imageTitle, &imageDescription, &imageCredit, &imageUrl, &imagePublished, &imageUpdated)
		if err != nil {
			fmt.Println("Error to create allImages:", err)
			return nil, err
		}

		imageDetails := ImageNew(imageId, imageTitle, imageDescription, imageCredit, imageUrl, imagePublished, imageUpdated)
		allImages = append(allImages, imageDetails)
	}

	return allImages, nil
}

func ImageEdit(getImage Image) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query("UPDATE images SET title = ?, description = ?, credit = ?, url = ?, published = ?, updated = ? WHERE id=?", getImage.Title, getImage.Description, getImage.Credit, getImage.Url, getImage.Published, getImage.Updated, getImage.Id)
	if err != nil {
		fmt.Println("Error on editing image:", err)
		return err
	}
	defer query.Close()

	return nil
}

func ImageDelete(getImageId int) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query("DELETE FROM images WHERE id=?", getImageId)
	if err != nil {
		fmt.Println("Error, not able to delete this image:", err)
		return err
	}
	defer query.Close()

	return nil
}

func ImageFindIt(getImageId int) ([]Image, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM images WHERE id=?", getImageId)
	if err != nil {
		fmt.Println("Error on image query:", err)
		return nil, err
	}
	defer rows.Close()

	var getImageData []Image
	for rows.Next() {
		var imageId int
		var imageTitle string
		var imageDescription string
		var imageCredit string
		var imageUrl string
		var imagePublished string
		var imageUpdated string
		err = rows.Scan(&imageId, &imageTitle, &imageDescription, &imageCredit, &imageUrl, &imagePublished, &imageUpdated)
		if err != nil {
			fmt.Println("Error getting image data:", err)
			return nil, err
		}

		imageDetails := ImageNew(imageId, imageTitle, imageDescription, imageCredit, imageUrl, imagePublished, imageUpdated)
		getImageData = append(getImageData, imageDetails)
	}

	return getImageData, nil
}

func ImageFindByUrlReturnItsId(getImageUrl string) (int, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM images WHERE url=?", getImageUrl)
	if err != nil {
		fmt.Println("Error on image query:", err)
		return 0, err
	}
	defer rows.Close()

	var getImageData Image
	for rows.Next() {
		var imageId int
		var imageTitle string
		var imageDescription string
		var imageCredit string
		var imageUrl string
		var imagePublished string
		var imageUpdated string
		err = rows.Scan(&imageId, &imageTitle, &imageDescription, &imageCredit, &imageUrl, &imagePublished, &imageUpdated)
		if err != nil {
			fmt.Println("Error getting image data:", err)
			return 0, err
		}

		imageDetails := ImageNew(imageId, imageTitle, imageDescription, imageCredit, imageUrl, imagePublished, imageUpdated)
		getImageData = imageDetails
	}

	return getImageData.Id, nil
}
