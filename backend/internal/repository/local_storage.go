package repository

import (
	"PollSystem/internal/domain"

	"github.com/google/uuid"
)

type LocalStorage struct {
	pollsList   []domain.Poll
	polls       map[uuid.UUID]domain.Poll
	pollOptions map[uuid.UUID]domain.Option
}

func (l LocalStorage) CastVote(pollId uuid.UUID, vote domain.Vote) error {
	panic("unimplemented")
}

func (l LocalStorage) GetOption(pollId uuid.UUID, optionId uuid.UUID) (domain.Option, error) {
	return l.pollOptions[optionId], nil
}

func (l LocalStorage) AddPoll(poll domain.Poll) error {
	l.polls[poll.Id] = poll

	for index := range poll.Option {
		option := poll.Option[index]
		l.pollOptions[option.Id] = option
	}

	return nil
}

func (l LocalStorage) GetAll() ([]domain.Poll, error) {
	return l.pollsList, nil
}

func NewLocalStorage() PollRepository {
	return LocalStorage{
		pollsList:   []domain.Poll{},
		polls:       map[uuid.UUID]domain.Poll{},
		pollOptions: map[uuid.UUID]domain.Option{},
	}
}
