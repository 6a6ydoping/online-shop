package utils

import (
	"net/mail"
)

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
