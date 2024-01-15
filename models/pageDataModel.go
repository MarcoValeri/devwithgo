package models

import "time"

type PageData struct {
	PageTitle   string
	CurrentYear int
}

func NewPageData(getPageTitle string) PageData {
	setPageData := PageData{
		PageTitle:   getPageTitle,
		CurrentYear: time.Now().Year(),
	}
	return setPageData
}
