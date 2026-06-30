package repository

import(
	"election-voting-system/internal/models"

)

type MemoryVoterRepository struct {
	Election *models.Election
}

func NewMemoryVoterRepository(e *models.Election) *MemoryVoterRepository {
	return &MemoryVoterRepository{
		Election: e,
	}
}

func (r *MemoryVoterRepository) GetAll() []models.Voter {
	return r.Election.Voters
}