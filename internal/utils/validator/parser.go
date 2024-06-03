package validator

import (
	"fmt"
	"time"
)

// ParseTimeNowToDate function which converts time.Now to 02/01/2006 date format string
func ParseTimeNowToDate() string {
	year, month, day := time.Now().Date()

	formattedDate := fmt.Sprintf("%02d/%02d/%d", day, month, year)

	return formattedDate
}
