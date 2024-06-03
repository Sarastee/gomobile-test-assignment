package service

import "errors"

const (
	errMsgNoDataFound = "no data found by given parameters"
)

var (
	ErrNoDataFound = errors.New(errMsgNoDataFound) // ErrNoDataFound is Err No Data Found Error object
)
