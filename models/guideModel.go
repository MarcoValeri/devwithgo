package models

import (
	"devwithgo/database"
	"fmt"
)

type Guide struct {
	Id          int
	Title       string
	Description string
	Url         string
	Published   string
	Updated     string
	Content     string
}

func GuideNew(getGuideId int, getGuideTitle, getGuideDescription, getGuideUrl, getGuidePublished, getGuideUpdated, getGuideContent string) Guide {
	setNewGuide := Guide{
		Id:          getGuideId,
		Title:       getGuideTitle,
		Description: getGuideDescription,
		Url:         getGuideUrl,
		Published:   getGuidePublished,
		Updated:     getGuideUpdated,
		Content:     getGuideContent,
	}
	return setNewGuide
}

func GuideAddNewToDB(getGuide Guide) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query(
		"INSERT INTO guides (title, description, url, published, updated, content) VALUES (?, ?, ?, ?, ?, ?)",
		getGuide.Title, getGuide.Description, getGuide.Url, getGuide.Published, getGuide.Updated, getGuide.Content,
	)
	if err != nil {
		return fmt.Errorf("error adding a new guide: %w", err)
	}
	defer query.Close()

	return nil
}

func GuideShowGuides() ([]Guide, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM guides")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allGuides []Guide
	for rows.Next() {
		var guideId int
		var guideTitle string
		var guideDescription string
		var guideUrl string
		var guidePublished string
		var guideUpdated string
		var guideContent string
		err = rows.Scan(&guideId, &guideTitle, &guideDescription, &guideUrl, &guidePublished, &guideUpdated, &guideContent)
		if err != nil {
			return nil, err
		}

		guideDetails := GuideNew(guideId, guideTitle, guideDescription, guideUrl, guidePublished, guideUpdated, guideContent)
		allGuides = append(allGuides, guideDetails)
	}

	return allGuides, nil
}
