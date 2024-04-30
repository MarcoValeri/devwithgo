package models

import "devwithgo/database"

type SitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

func SitemapAllURLs() ([]SitemapURL, error) {

	var setURLsList []SitemapURL

	// Set URLs that are not stored in the db
	urlZero := SitemapURL{"https://www.devwithgo.dev/", "2024-04-28"}
	urlOne := SitemapURL{"https://www.devwithgo.dev/guides/guides", "2024-04-28"}
	setURLsList = append(setURLsList, urlZero)
	setURLsList = append(setURLsList, urlOne)

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

	return setURLsList, nil
}
