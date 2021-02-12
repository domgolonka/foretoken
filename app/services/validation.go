package services

import (
	"fmt"
	"strings"
)

var ErrMissing = "MISSING"
var ErrTaken = "TAKEN"
var ErrFormatInvalid = "FORMAT_INVALID"
var ErrLocked = "LOCKED"
var ErrExpired = "EXPIRED"
var ErrNotFound = "NOT_FOUND"

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e FieldError) String() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

func (e FieldError) Error() string {
	return e.String()
}

type FieldErrors []FieldError

func (es FieldErrors) Error() string {
	var buf = make([]string, len(es))
	for i, e := range es {
		buf[i] = e.Error()
	}
	return strings.Join(buf, ", ")
}
