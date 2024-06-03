package validator

import (
	"errors"
	"net/http"
	"time"
)

// CheckContentType function which validate request content type
func CheckContentType(r *http.Request) error {
	if len(r.Header.Values("Content-Type")) == 0 {
		return errors.New("empty content-type")
	}

	for contentTypeCurrentIndex, contentType := range r.Header.Values("Content-type") {
		if contentType == "application/json" {
			break
		}

		if contentTypeCurrentIndex == len(r.Header.Values("Content-Type"))-1 {
			return errors.New("wrong content-type")
		}
	}

	return nil
}

// ValidateDateQueryParam function which validates Date query param
func ValidateDateQueryParam(input string) bool {
	_, err := time.Parse("02/01/2006", input)
	return err == nil
}
