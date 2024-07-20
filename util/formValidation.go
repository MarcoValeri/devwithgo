package util

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"net/mail"
	"path/filepath"
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

func FormIsValidImage(file multipart.File, fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	if !allowedExts[ext] {
		fmt.Println("Error, invalid image type")
		return false
	}

	// Check MIME
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading the file")
		return false
	}

	contentType := http.DetectContentType(buffer)
	if !strings.HasPrefix(contentType, "image/") {
		fmt.Println("Error, invalid image content")
		return false
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error resetting file position")
		return false
	}

	return true
}
