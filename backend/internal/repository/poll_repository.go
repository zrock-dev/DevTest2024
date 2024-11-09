package repository

import (
	"PollSystem/internal/domain"

	"github.com/google/uuid"
)


type PollRepository interface {
    GetAll() []domain.Poll
    AddPoll(poll domain.Poll)
    CastVote(pollId uuid.UUID, vote domain.Vote)
    GetAllOptions(pollId uuid.UUID) []domain.Option
}
