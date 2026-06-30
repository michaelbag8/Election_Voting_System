package repository

import "election-voting-system/internal/models"

type VoterRepository interface {
	GetAll() []models.Voter
	Create(models.Voter) error
	FindByVoterID(string) (*models.Voter, error)
}