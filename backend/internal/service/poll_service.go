package service

import (
	"PollSystem/internal/domain"

	"github.com/google/uuid"
)

type PollService interface {
	GetAllPoll() ([]domain.PollResponse, error)
	CastVote(pollId uuid.UUID, vote domain.PollVoteRequest) error
	CreatePoll(domain.PollRequest) (domain.Poll, error)
}
