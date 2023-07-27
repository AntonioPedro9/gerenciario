package utils

import "time"

func GetTodaysDate() string {
	dateFormat := "01/02/2003"
	formatedDate := time.Now().Format(dateFormat)

	return formatedDate
}
