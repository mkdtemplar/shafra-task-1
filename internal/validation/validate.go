package validation

import (
	"errors"
	"regexp"
)

var (
	isValidFullName = regexp.MustCompile("^[a-zA-Z\\s]+$").MatchString
	isValidAge      = regexp.MustCompile("^(\\d+)$").MatchString
)

type Validation struct {
	Err error
}

func (v *Validation) Error() string {
	var err string
	return err
}

func (v *Validation) ValidateFullName(value string) *Validation {
	if !isValidFullName(value) {
		v.Err = errors.New("invalid full name format, only letters")
		return v
	}
	return v
}

func (v *Validation) ValidateAge(value string) *Validation {
	if !isValidAge(value) {
		v.Err = errors.New("invalid age format, only integers")
		return v
	}

	return v
}
