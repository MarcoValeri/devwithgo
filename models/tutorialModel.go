package models

import (
	"devwithgo/database"
	"fmt"
)

type Tutorial struct {
	Id          int
	Title       string
	Description string
	Url         string
	Published   string
	Updated     string
	ImageId     int
	Content     string
}

type TutorialWithRelatedImage struct {
	Id          int
	Title       string
	Description string
	Url         string
	Published   string
	Updated     string
	ImageId     int
	ImageUrl    string
	ImageAlt    string
	Content     string
}

func TutorialNew(getTutorialId int, getTutorialTitle string, getTutorialDescription string, getTutorialUrl string, getTutorialPublished string, getTutorialUpdated string, getTutorialImageId int, getTutorialContent string) Tutorial {
	setNewTutorial := Tutorial{
		Id:          getTutorialId,
		Title:       getTutorialTitle,
		Description: getTutorialDescription,
		Url:         getTutorialUrl,
		Published:   getTutorialPublished,
		Updated:     getTutorialUpdated,
		ImageId:     getTutorialImageId,
		Content:     getTutorialContent,
	}
	return setNewTutorial
}

func TutorialWithRelatedImageNew(getTutorialId int, getTutorialTitle string, getTutorialDescription string, getTutorialUrl string, getTutorialPublished string, getTutorialUpdated string, getTutorialImageId int, getTutorialImageUrl string, getTutorialImageAlt string, getTutorialContent string) TutorialWithRelatedImage {
	setNewTutorialWithRelatedImage := TutorialWithRelatedImage{
		Id:          getTutorialId,
		Title:       getTutorialTitle,
		Description: getTutorialDescription,
		Url:         getTutorialUrl,
		Published:   getTutorialPublished,
		Updated:     getTutorialUpdated,
		ImageId:     getTutorialImageId,
		ImageUrl:    getTutorialImageUrl,
		ImageAlt:    getTutorialImageAlt,
		Content:     getTutorialContent,
	}
	return setNewTutorialWithRelatedImage
}

func TutorialAddNewToDB(getTutorial Tutorial) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query(
		"INSERT INTO tutorials (title, description, url, published, updated, image_id, content) VALUES (?, ?, ?, ?, ?, ?, ?)",
		getTutorial.Title, getTutorial.Description, getTutorial.Url, getTutorial.Published, getTutorial.Updated, getTutorial.ImageId, getTutorial.Content,
	)
	if err != nil {
		fmt.Println("Error adding a new tutorial:", err)
		return err
	}
	defer query.Close()

	return nil
}

func TutorialShowTutorials() ([]TutorialWithRelatedImage, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT tutorials.id, tutorials.title, tutorials.description, tutorials.url, tutorials.published, tutorials.updated, tutorials.image_id, images.url, images.description, tutorials.content FROM tutorials JOIN images ON tutorials.image_id = images.id")
	if err != nil {
		fmt.Println("Error getting tutorials from the db:", err)
		return nil, err
	}
	defer rows.Close()

	var allTutorials []TutorialWithRelatedImage
	for rows.Next() {
		var tutorialId int
		var tutorialTitle string
		var tutotialDescription string
		var tutorialUrl string
		var tutorialPublished string
		var tutorialUpdated string
		var tutorialImageId int
		var tutorialImageUrl string
		var tutorialImageAlt string
		var tutorialContent string
		err = rows.Scan(&tutorialId, &tutorialTitle, &tutotialDescription, &tutorialUrl, &tutorialPublished, &tutorialUpdated, &tutorialImageId, &tutorialImageUrl, &tutorialImageAlt, &tutorialContent)
		if err != nil {
			return nil, err
		}

		tutorialDetails := TutorialWithRelatedImageNew(tutorialId, tutorialTitle, tutotialDescription, tutorialUrl, tutorialPublished, tutorialUpdated, tutorialImageId, tutorialImageUrl, tutorialImageAlt, tutorialContent)
		allTutorials = append(allTutorials, tutorialDetails)
	}

	return allTutorials, nil
}

func TutorialGetPublishedTutorials() ([]TutorialWithRelatedImage, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT tutorials.id, tutorials.title, tutorials.description, tutorials.url, tutorials.published, tutorials.updated, tutorials.image_id, images.url, images.description, tutorials.content FROM tutorials JOIN images ON tutorials.image_id = images.id WHERE tutorials.published < NOW() ORDER BY tutorials.published ASC")
	if err != nil {
		fmt.Println("Error on getting pulished tutorials:", err)
		return nil, err
	}
	defer rows.Close()

	var allTutorials []TutorialWithRelatedImage
	for rows.Next() {
		var tutorialId int
		var tutorialTitle string
		var tutotialDescription string
		var tutorialUrl string
		var tutorialPublished string
		var tutorialUpdated string
		var tutorialImageId int
		var tutorialImageUrl string
		var tutorialImageAlt string
		var tutorialContent string
		err = rows.Scan(&tutorialId, &tutorialTitle, &tutotialDescription, &tutorialUrl, &tutorialPublished, &tutorialUpdated, &tutorialImageId, &tutorialImageUrl, &tutorialImageAlt, &tutorialContent)
		if err != nil {
			return nil, err
		}

		tutorialDetail := TutorialWithRelatedImageNew(tutorialId, tutorialTitle, tutotialDescription, tutorialUrl, tutorialPublished, tutorialUpdated, tutorialImageId, tutorialImageUrl, tutorialImageAlt, tutorialContent)
		allTutorials = append(allTutorials, tutorialDetail)
	}

	return allTutorials, nil
}
