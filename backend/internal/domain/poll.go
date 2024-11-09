package domain

import "github.com/google/uuid"

type Vote struct {
	Id       uuid.UUID
	OptionId uuid.UUID
	Email    string
}

type Option struct {
	Id   uuid.UUID
	Name string
	Vote []Vote
}

type Poll struct {
	Id     uuid.UUID
	Name   string
	Option []Option
}
