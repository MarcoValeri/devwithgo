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
	Image       string
	Content     string
}

func TutorialNew(getTutorialId int, getTutorialTitle, getTutorialDescription, getTutorialUrl, getTutorialPublished, getTutorialUpdated, getTutorialImage, getTutorialContent string) Tutorial {
	setNewTutorial := Tutorial{
		Id:          getTutorialId,
		Title:       getTutorialTitle,
		Description: getTutorialDescription,
		Url:         getTutorialUrl,
		Published:   getTutorialPublished,
		Updated:     getTutorialUpdated,
		Image:       getTutorialImage,
		Content:     getTutorialContent,
	}
	return setNewTutorial
}

func TutorialShowTutorials() ([]Tutorial, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tutorials")
	if err != nil {
		fmt.Println("Error getting tutorials from the db:", err)
		return nil, err
	}
	defer rows.Close()

	var allTutorials []Tutorial
	for rows.Next() {
		var tutorialId int
		var tutorialTitle string
		var tutotialDescription string
		var tutorialUrl string
		var tutorialPublished string
		var tutorialUpdated string
		var tutorialImage string
		var tutorialContent string
		err = rows.Scan(&tutorialId, &tutorialTitle, &tutotialDescription, &tutorialUrl, &tutorialPublished, &tutorialUpdated, &tutorialImage, &tutorialContent)
		if err != nil {
			return nil, err
		}

		tutorialDetails := TutorialNew(tutorialId, tutorialTitle, tutotialDescription, tutorialUrl, tutorialPublished, tutorialUpdated, tutorialImage, tutorialContent)
		allTutorials = append(allTutorials, tutorialDetails)
	}

	return allTutorials, nil
}
