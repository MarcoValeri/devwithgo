package util

import (
	"net/mail"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

func FormEmailInput(getEmailInput string) bool {
	/*
	* Get
	* @param string getEmilInput
	* and
	* @return bool true if the email is valid format
	* false otherwise
	 */
	_, err := mail.ParseAddress(getEmailInput)
	return err == nil
}

func FormEmailLengthInput(getEmailInput string) bool {
	if len(getEmailInput) < 5 || len(getEmailInput) > 40 {
		return false
	}
	return true
}

func FormPasswordInput(getPassword string) bool {
	/**
	* Password shoul:
	* 	1 - longer than 8 charactes
	*	2 - no longer than 20 charactes
	 */
	if len(getPassword) < 8 || len(getPassword) > 20 {
		return false
	}
	return true
}

func FormSanitizeStringInput(getStringInput string) string {
	/**
	* Avoid HTML injection
	* Remove space at the left and right position
	 */
	sanitizeHtml := bluemonday.StrictPolicy()
	outputInput := sanitizeHtml.Sanitize(getStringInput)

	outputInput = strings.TrimSpace(outputInput)
	return outputInput
}
