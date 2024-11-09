package domain

import "github.com/google/uuid"


type OptionResponse struct {
	Id   uuid.UUID
	Name string
	AmountVotes int
}

type PollResponse struct {
	Id uuid.UUID
	Name   string
	Option []OptionResponse
}


