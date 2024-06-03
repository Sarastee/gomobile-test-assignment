package api

import "errors"

const (
	errMsgInvalidDateFormat = "invalid date format"
	errMsgInvalidValFormat  = "invalid valute format"
)

var (
	ErrInvalidDateFormat = errors.New(errMsgInvalidDateFormat) // ErrInvalidDateFormat is Err Invalid Date Format Error object
	ErrInvalidValFormat  = errors.New(errMsgInvalidValFormat)  // ErrInvalidValFormat is Err Invalid Val Format Error object
)
