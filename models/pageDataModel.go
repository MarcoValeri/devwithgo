package models

import "time"

type PageData struct {
	PageTitle       string
	PageDescription string
	CurrentYear     int
}

func NewPageData(getPageTitle, getPageDescription string) PageData {
	setPageData := PageData{
		PageTitle:       getPageTitle,
		PageDescription: getPageDescription,
		CurrentYear:     time.Now().Year(),
	}
	return setPageData
}
