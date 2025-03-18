package repositories

import (
	"rvkc/config"
	"rvkc/models"
)

type PilotRepository interface {
	CreatePilot(pilot *models.Pilot) error
	GetPilots() ([]models.Pilot, error)
	GetPilotByDocument(document string) (*models.Pilot, error)
	UpdatePilot(pilot *models.Pilot) error
	DeletePilot(id uint) error
}

type pilotRepository struct{}

func NewPilotRepository() PilotRepository {
	return &pilotRepository{}
}

func (r *pilotRepository) CreatePilot(pilot *models.Pilot) error {
	return config.DB.Create(pilot).Error
}

func (r *pilotRepository) GetPilots() ([]models.Pilot, error) {
	var pilots []models.Pilot
	err := config.DB.Find(&pilots).Error
	return pilots, err
}

func (r *pilotRepository) GetPilotByDocument(document string) (*models.Pilot, error) {
	var pilot models.Pilot
    err := config.DB.Where("document = ?", document).First(&pilot).Error
    return &pilot, err
}

func (r *pilotRepository) UpdatePilot(pilot *models.Pilot) error {
    return config.DB.Save(pilot).Error
}


func (r *pilotRepository) DeletePilot(id uint) error {
	return config.DB.Delete(&models.Pilot{}, id).Error
}
