package trait

import (
	"fmt"
	"hastane-takip/internal/db"
	"hastane-takip/internal/models"
	"time"

	"gorm.io/gorm"
)

type IStaffHandler interface {
	CreateStaff(staff *models.Staff) error
	UpdateStaff(id uint, updatedStaff *models.Staff) error
	DeleteStaff(id uint) error
	IsExist(staffID uint) (bool, error)
	GetStaffByHospitalID(hospitalID uint) ([]models.Staff, error)
	GetHospitalIDByUserID(userID uint) (uint, error)
	GetStaffByID(id uint) (*models.Staff, error)
	GetJobCategoriesFromDB() ([]models.JobCategory, error)
}

type StaffHandler struct {
	DB *gorm.DB
}

func NewStaffHandler() IStaffHandler {
	return &StaffHandler{
		DB: db.GetDB(),
	}
}

func (s *StaffHandler) CreateStaff(staff *models.Staff) error {
	result := s.DB.Exec("INSERT INTO staffs (first_name, last_name, email, phone_number, role) VALUES (?, ?, ?, ?, ?)", staff.FirstName, staff.LastName, staff.Email, staff.PhoneNumber, staff.Role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *StaffHandler) UpdateStaff(id uint, staffUpdate *models.Staff) error {
	query := `UPDATE staffs
              SET first_name = ?, last_name = ?, email = ?, phone_number = ?, role = ?, updated_at = ?
              WHERE id = ?`
	result := s.DB.Exec(query, staffUpdate.FirstName, staffUpdate.LastName, staffUpdate.Email, staffUpdate.PhoneNumber, staffUpdate.Role, time.Now(), id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("staff record not found")
	}

	return nil
}

func (s *StaffHandler) DeleteStaff(id uint) error {
	query := `DELETE FROM staffs WHERE id = ?`
	result := s.DB.Exec(query, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("staff record not found")
	}
	return nil
}

func (s *StaffHandler) IsExist(staffID uint) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM staffs WHERE id = ?)`
	result := s.DB.Raw(query, staffID).Scan(&exists)
	if result.Error != nil {
		return false, result.Error
	}

	return exists, nil
}

func (h *StaffHandler) GetHospitalIDByUserID(userID uint) (uint, error) {
	var hospitalID uint
	query := `SELECT hospital_id FROM users WHERE id = ?`
	err := h.DB.Raw(query, userID).Scan(&hospitalID).Error
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve hospital ID: %w", err)
	}
	return hospitalID, nil
}

func (h *StaffHandler) GetStaffByHospitalID(hospitalID uint) ([]models.Staff, error) {
	var staffList []models.Staff
	query := `SELECT * FROM staffs WHERE hospital_id = ?`
	err := h.DB.Raw(query, hospitalID).Find(&staffList).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve staff list: %w", err)
	}
	return staffList, nil
}

func (s *StaffHandler) GetStaffByID(id uint) (*models.Staff, error) {
	var staff models.Staff
	if err := s.DB.Raw("SELECT * FROM staffs WHERE id = ?", id).Scan(&staff).Error; err != nil {
		return nil, fmt.Errorf("err")
	}
	return &staff, nil
}

func (s *StaffHandler) GetJobCategoriesFromDB() ([]models.JobCategory, error) {
	var categories []models.JobCategory
	if err := s.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
