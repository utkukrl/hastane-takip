package userroutes

import (
	"hastane-takip/internal/db"
	"hastane-takip/internal/models"
	"hastane-takip/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RegisterRoute struct {
	DB *gorm.DB
}

func NewRegisterRoute() *RegisterRoute {
	return &RegisterRoute{
		DB: db.GetDB(),
	}
}

func (r *RegisterRoute) Handler(c *fiber.Ctx) error {
	var request struct {
		FirstName   string          `json:"first_name"`
		LastName    string          `json:"last_name"`
		Email       string          `json:"email"`
		PhoneNumber string          `json:"phone_number"`
		Role        string          `json:"role"`
		Password    string          `json:"password"`
		Hospital    models.Hospital `json:"hospital"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}
	hospital := models.Hospital{
		Name:        request.Hospital.Name,
		Address:     request.Hospital.Address,
		City:        request.Hospital.City,
		State:       request.Hospital.State,
		PostalCode:  request.Hospital.PostalCode,
		PhoneNumber: request.Hospital.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := `INSERT INTO hospitals (name, address, city, state, postal_code, phone_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	if err := r.DB.Raw(query, hospital.Name, hospital.Address, hospital.City, hospital.State, hospital.PostalCode, hospital.PhoneNumber, hospital.CreatedAt, hospital.UpdatedAt).Scan(&hospital.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create hospital",
		})
	}

	staff := models.Staff{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Role:        request.Role,
		Is_Admin:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Password:    hashedPassword,
	}

	query = `INSERT INTO staffs (first_name, last_name, email, phone_number, role, is_admin, created_at, updated_at, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	if err := r.DB.Raw(query, staff.FirstName, staff.LastName, staff.Email, staff.PhoneNumber, staff.Role, staff.Is_Admin, staff.CreatedAt, staff.UpdatedAt, staff.Password).Scan(&staff.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create staff",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "User and hospital created successfully",
		"staff":    staff,
		"hospital": hospital,
	})
}
