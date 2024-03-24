package util

import (
	"fmt"
	"time"
)

func DateContentValidation(getArticleDate string) bool {
	setArticleDate, err := time.Parse("2006-01-02 15:04:05", getArticleDate)
	if err != nil {
		fmt.Println("Error:", err)
	}

	getDateNow := time.Now()

	return setArticleDate.Before(getDateNow)
}
