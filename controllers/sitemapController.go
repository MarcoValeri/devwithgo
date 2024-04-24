package controllers

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type SitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

type Sitemap struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	URLs    []SitemapURL `xml:"url"`
}

func SitemapController() {
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")

		urls := []SitemapURL{
			{"https://www.devwithgo.dev/", "2024-04-24"},
			{"https://www.devwithgo.dev/about", "2024-04-24"},
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
