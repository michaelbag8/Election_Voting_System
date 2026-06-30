package services

import (
	"election-voting-system/internal/models"
	"errors"
	"strings"
)

type VoterService struct {
	Election *models.Election
}

func NewVoterService(e *models.Election) *VoterService {
	return &VoterService{Election: e}
}

func (s *VoterService) CreateVoter(name, voterID string) (models.Voter, error) {

	name = strings.TrimSpace(name)
	voterID = strings.TrimSpace(voterID)

	if name == "" || voterID == "" {
		return models.Voter{}, errors.New("name and voterID cannot be empty")
	}

	// check duplicates
	for _, v := range s.Election.Voters {
		if v.VoterID == voterID {
			return models.Voter{}, errors.New("voter already exists")
		}
	}

	voter := models.Voter{
		ID:       len(s.Election.Voters) + 1,
		Name:     name,
		VoterID:  voterID,
		HasVoted: false,
	}

	s.Election.Voters = append(s.Election.Voters, voter)

	return voter, nil
}