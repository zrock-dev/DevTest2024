package service

import "net/mail"

type PollValidator struct{}

// TODO: test use cases of different emails
func checkEmailFormat(email string) error {
	_, err  := mail.ParseAddress(email)
	return err
}
