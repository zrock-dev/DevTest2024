package repository

import (
	"PollSystem/internal/domain"

	"github.com/google/uuid"
)

type LocalStorage struct {
	polls       []domain.Poll
	pollOptions map[uuid.UUID][]domain.Option
	optionVotes map[uuid.UUID][]domain.Vote
}

func (l LocalStorage) GetAllOptions(pollId uuid.UUID) []domain.Option {
    return l.pollOptions[pollId]
}

func (l LocalStorage) AddOption(pollId uuid.UUID, option domain.Option) {
	l.pollOptions[pollId] =  append(l.pollOptions[pollId], option)
}

func (l LocalStorage) AddPoll(poll domain.Poll) {
	l.polls = append(l.polls, poll)
}

func (l LocalStorage) CastVote(pollId uuid.UUID, vote domain.Vote) {
	l.optionVotes[vote.OptionId] =  append(l.optionVotes[pollId], vote)
}

func (l LocalStorage) GetAll() []domain.Poll {
    return l.polls
}

func NewLocalStorage() PollRepository {
	return LocalStorage{
		polls:       []domain.Poll{},
		pollOptions: map[uuid.UUID][]domain.Option{},
		optionVotes: map[uuid.UUID][]domain.Vote{},
	}
}
