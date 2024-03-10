package models

type Guide struct {
	Id          int
	Title       string
	Description string
	Url         string
	Published   string
	Updated     string
	Content     string
}

func GuideNew(getGuide Guide) Guide {
	setNewGuide := getGuide
	return setNewGuide
}
