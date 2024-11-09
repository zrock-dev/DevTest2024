package repository

import (
	"PollSystem/internal/domain"

	"github.com/google/uuid"
)


type PollRepository interface {
    GetAll() ([]domain.Poll, error)
    AddPoll(poll domain.Poll) error
    CastVote(pollId uuid.UUID, vote domain.Vote) error
    GetOption(pollId uuid.UUID, optionId uuid.UUID) (domain.Option, error)
}
