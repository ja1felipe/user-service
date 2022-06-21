package models

import (
	"errors"
	"regexp"
)

func BasicUserValidation(u *User) (err error) {

	if u.Identification != "" {
		validIdentification, _ := regexp.MatchString("(([0-9]{3}.[0-9]{3}.[0-9]{3}-[0-9]{2})|([0-9]{11}))", u.Identification)

		if !validIdentification {
			return errors.New("invalidated Identification number")
		}
	}

	validEmail, _ := regexp.MatchString("\\S+@\\S+\\.\\S+", u.Email)

	if !validEmail {
		return errors.New("invalidated email")
	}
	return
}
