package controllers

import (
	"devwithgo/models"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Sitemap struct {
	XMLName xml.Name            `xml:"urlset"`
	Xmlns   string              `xml:"xmlns,attr"`
	URLs    []models.SitemapURL `xml:"url"`
}

func SitemapController() {
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")

		urls, err := models.SitemapAllURLs()
		if err != nil {
			fmt.Println("Error:", err)
		}

		sitemap := Sitemap{
			Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
			URLs:  urls,
		}

		output, err := xml.MarshalIndent(sitemap, "", " ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, xml.Header+string(output))
	})
}
