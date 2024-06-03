package repository

import "errors"

const (
	errMsgCacheNotFound = "data not found in cache by provided key"
)

var ErrCacheNotFound = errors.New(errMsgCacheNotFound) // ErrCacheNotFound is Err Cache Not Found Error Object
