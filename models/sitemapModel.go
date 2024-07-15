package models

import (
	"devwithgo/database"
	"fmt"
)

type SitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

func SitemapAllURLs() ([]SitemapURL, error) {

	var setURLsList []SitemapURL

	// Set URLs that are not stored in the db
	urlZero := SitemapURL{"https://www.devwithgo.dev/", "2024-04-28"}
	urlOne := SitemapURL{"https://www.devwithgo.dev/guides/guides", "2024-04-28"}
	urlThree := SitemapURL{"https://www.devwithgo.dev/tutorials/tutorials", "2024-08-01"}
	setURLsList = append(setURLsList, urlZero)
	setURLsList = append(setURLsList, urlOne)
	setURLsList = append(setURLsList, urlThree)

	// Get all guides URLs
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT url, updated FROM guides WHERE published < NOW()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urlGuide SitemapURL
	for rows.Next() {
		var guideUrl string
		var guideUpdated string
		err = rows.Scan(&guideUrl, &guideUpdated)
		if err != nil {
			return nil, err
		}
		urlGuide.Loc = "https://www.devwithgo.dev/guide/" + guideUrl
		urlGuide.LastMod = guideUpdated[:10]
		setURLsList = append(setURLsList, urlGuide)
	}

	// Get all tutorials
	rowsTutorial, errTutorial := db.Query("SELECT url, updated FROM tutorials WHERE published < NOW()")
	if errTutorial != nil {
		fmt.Println("Error to query tutorials for sitemapModel:", errTutorial)
		return nil, errTutorial
	}
	defer rowsTutorial.Close()

	var urlTutorial SitemapURL
	for rowsTutorial.Next() {
		var tutorialUrl string
		var tutorialUpdated string
		errTutorial = rowsTutorial.Scan(&tutorialUrl, &tutorialUpdated)
		if errTutorial != nil {
			fmt.Println("Error saving tutorial data for the sitemap:", errTutorial)
			return nil, errTutorial
		}
		urlTutorial.Loc = "https://www.devwithgo.dev/tutorial/" + tutorialUrl
		urlTutorial.LastMod = tutorialUpdated[:10]
		setURLsList = append(setURLsList, urlTutorial)
	}

	// Get all the images
	rowsImage, errImage := db.Query("SELECT url, updated FROM images WHERE published < NOW()")
	if errImage != nil {
		fmt.Println("Error to query images for sitemapModel:", errImage)
		return nil, errImage
	}
	defer rowsImage.Close()

	var urlImage SitemapURL
	for rowsImage.Next() {
		var imageUrl string
		var imageUpdated string
		errImage = rowsImage.Scan(&imageUrl, &imageUpdated)
		if errImage != nil {
			fmt.Println("Error saveing image data firn the sitemap:", errImage)
			return nil, errImage
		}
		urlImage.Loc = "https://www.devwithgo.dev/public/images/" + imageUrl
		urlImage.LastMod = imageUpdated[:10]
		setURLsList = append(setURLsList, urlImage)
	}

	return setURLsList, nil
}
