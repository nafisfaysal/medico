package models

import "errors"

var ErrUnknown = errors.New("unknown error")
var ErrRecordNotFound = errors.New("record not found")
var ErrNotAuthorized = errors.New("not authorized")
var ErrInvalidUserOrPassword = errors.New("invalid user or password")
