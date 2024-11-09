package domain

import "github.com/google/uuid"

type OptionRequest struct {
	Name string
}

type PollRequest struct {
	Name string
	Options []OptionRequest
}

type PollVoteRequest struct {
	OptionId uuid.UUID
	EmailAddress string
}
