package services

import (
	"rvkc/models"
	"rvkc/repositories"
)

type PilotService interface {
	CreatePilot(pilot *models.Pilot) error
	GetPilots() ([]models.Pilot, error)
	GetPilotByDocument(document string) (*models.Pilot, error)
	UpdatePilot(pilot *models.Pilot) error
	DeletePilot(id uint) error
}

type pilotService struct {
	repository repositories.PilotRepository
}

func NewPilotService(repo repositories.PilotRepository) PilotService {
	return &pilotService{repository: repo}
}

func (s *pilotService) CreatePilot(pilot *models.Pilot) error {
	return s.repository.CreatePilot(pilot)
}

func (s *pilotService) GetPilots() ([]models.Pilot, error) {
	return s.repository.GetPilots()
}

func (s *pilotService) GetPilotByDocument(document string) (*models.Pilot, error) {
	return s.repository.GetPilotByDocument(document)
}

func (s *pilotService) UpdatePilot(pilot *models.Pilot) error {
	return s.repository.UpdatePilot(pilot)
}

func (s *pilotService) DeletePilot(id uint) error {
	return s.repository.DeletePilot(id)
}
