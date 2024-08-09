package trait

import (
	"fmt"
	"hastane-takip/internal/db"
	"hastane-takip/internal/models"

	"gorm.io/gorm"
)

type IHospitalHandler interface {
	GetHospitalByID(id uint) (*models.Hospital, error)
	GetProvincesFromDB() ([]models.Province, error)
}

type HospitalHandler struct {
	DB *gorm.DB
}

func NewHospitalHandler() IHospitalHandler {
	return &HospitalHandler{
		DB: db.GetDB(),
	}
}

func (h *HospitalHandler) GetHospitalByID(userID uint) (*models.Hospital, error) {
	var hospitalID uint
	staffQuery := `SELECT hospital_id FROM staffs WHERE user_id = ?`
	err := h.DB.Raw(staffQuery, userID).Scan(&hospitalID).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve hospital ID from staffs: %w", err)
	}
	if hospitalID == 0 {
		return nil, fmt.Errorf("no hospital found for user ID %d", userID)
	}
	var hospital models.Hospital

	hospitalQuery := `SELECT * FROM hospitals WHERE id = ?`
	err = h.DB.Raw(hospitalQuery, hospitalID).Scan(&hospital).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve hospital by ID: %w", err)
	}

	return &hospital, nil
}

func (s *HospitalHandler) GetProvincesFromDB() ([]models.Province, error) {
	var provinces []models.Province
	query := `SELECT * FROM provinces`
	err := s.DB.Raw(query).Scan(&provinces).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve provinces: %w", err)
	}
	return provinces, nil
}
