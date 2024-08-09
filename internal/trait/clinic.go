package trait

import (
	"fmt"
	"hastane-takip/internal/db"
	"hastane-takip/internal/models"
	"time"

	"gorm.io/gorm"
)

type IClinicHandler interface {
	GetClinicsByHospitalID(hospitalID uint) ([]models.Clinic, error)
	GetClinicByID(id uint) (*models.Clinic, error)
	UpdateClinic(id uint, clinicUpdate *models.Clinic) error
	DeleteClinic(id uint) error
	CreateClinic(clinic *models.Clinic) error
}

type ClinicHandler struct {
	DB *gorm.DB
}

func NewClinicHandler() IClinicHandler {
	return &ClinicHandler{
		DB: db.GetDB(),
	}
}

func (h *ClinicHandler) GetClinicsByHospitalID(hospitalID uint) ([]models.Clinic, error) {
	var clinicList []models.Clinic
	query := `SELECT * FROM clinics WHERE hospital_id = ?`
	err := h.DB.Raw(query, hospitalID).Scan(&clinicList).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve clinic list: %w", err)
	}
	return clinicList, nil
}

func (h *ClinicHandler) GetClinicByID(id uint) (*models.Clinic, error) {
	var clinic models.Clinic
	err := h.DB.First(&clinic, id).Error
	if err != nil {
		return nil, fmt.Errorf("clinic record not found: %w", err)
	}
	return &clinic, nil
}

func (h *ClinicHandler) UpdateClinic(id uint, clinicUpdate *models.Clinic) error {
	var existingClinic models.Clinic
	err := h.DB.First(&existingClinic, id).Error
	if err != nil {
		return fmt.Errorf("clinic record not found: %w", err)
	}

	if clinicUpdate.Name != "" {
		existingClinic.Name = clinicUpdate.Name
	}
	if clinicUpdate.Address != "" {
		existingClinic.Address = clinicUpdate.Address
	}
	if clinicUpdate.PhoneNumber != "" {
		existingClinic.PhoneNumber = clinicUpdate.PhoneNumber
	}

	existingClinic.UpdatedAt = time.Now()
	err = h.DB.Save(&existingClinic).Error
	if err != nil {
		return fmt.Errorf("failed to update clinic record: %w", err)
	}

	return nil
}

func (h *ClinicHandler) DeleteClinic(id uint) error {
	query := `DELETE FROM clinics WHERE id = ?`
	result := h.DB.Exec(query, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("clinic record not found")
	}
	return nil
}

func (h *ClinicHandler) CreateClinic(clinic *models.Clinic) error {
	query := `INSERT INTO clinics (name, address, city, state, postal_code, phone_number) VALUES (?, ?, ?, ?, ?, ?)`
	result := h.DB.Exec(query, clinic.Name, clinic.Address, clinic.City, clinic.State, clinic.PostalCode, clinic.PhoneNumber)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
