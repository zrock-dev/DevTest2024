package service

import (
	"PollSystem/internal/domain"
	"PollSystem/internal/repository"

	"github.com/google/uuid"
)

type LocalStorageService struct {
	repository repository.PollRepository
}

func (l LocalStorageService) CastVote(pollId uuid.UUID, vote domain.PollVoteRequest) error {
	if err := checkEmailFormat(vote.EmailAddress); err != nil {
		return err
	}
	
	option, err := l.repository.GetOption(pollId, vote.OptionId)
	if err != nil {
		return err
	}

	option.Vote = append(option.Vote, domain.Vote{
		Id:       uuid.New(),
		Email:    vote.EmailAddress,
	})
	
	return nil
}

func (l LocalStorageService) CreatePoll(pollRequest domain.PollRequest) (domain.Poll, error) {
	var options []domain.Option
	for index := range pollRequest.Options {
		optionRequest := pollRequest.Options[index]
		options = append(options, domain.Option{
			Id:   uuid.New(),
			Name: optionRequest.Name,
			Vote: []domain.Vote{},
		})
	}

	poll := domain.Poll {
		Id:     uuid.New(),
		Name:   pollRequest.Name,
		Option: []domain.Option{},
	}

	return poll, l.repository.AddPoll(poll)

}

func handleOptions(pollOptions []domain.Option) []domain.OptionResponse {
	var responses []domain.OptionResponse

	for index := range pollOptions {
		option := pollOptions[index]
		response := domain.OptionResponse{
			Id:          option.Id,
			Name:        option.Name,
			AmountVotes: len(option.Vote),
		}
		responses = append(responses, response)
	}

	return responses
}

func (l LocalStorageService) GetAllPoll() ([]domain.PollResponse, error) {
	polls, err := l.repository.GetAll()
	if err != nil {
		return nil, err
	}

	var pollsResponse []domain.PollResponse
	for index := range polls {
		poll := polls[index]
		response := domain.PollResponse{
			Id:     poll.Id,
			Name:   poll.Name,
			Option: handleOptions(poll.Option),
		}

		pollsResponse = append(pollsResponse, response)
	}

	return pollsResponse, nil
}

func NewLocalStorageService(repository repository.PollRepository) PollService {
	return LocalStorageService{
		repository: repository,
	}
}
