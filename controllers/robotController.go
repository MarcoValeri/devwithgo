package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

func RobotController() {
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		output := []string{
			"User-agent: *",
			"Sitemap: https://www.devwithgo.dev/sitemap.xml",
			"Disallow: /admin/",
		}

		fmt.Fprint(w, strings.Join(output, "\n"))
	})
}
