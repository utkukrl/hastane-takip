package userroutes

import (
	"hastane-takip/internal/db"
	"hastane-takip/internal/models"
	"hastane-takip/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginRoute struct {
	DB *gorm.DB
}

func NewLoginRoute() *LoginRoute {
	return &LoginRoute{
		DB: db.GetDB(),
	}
}
func (r *LoginRoute) Handler(c *fiber.Ctx) error {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var staff models.Staff
	query := `SELECT id, first_name, last_name, email, phone_number, role, clinic, is_admin, password FROM staffs WHERE email = $1`
	if err := r.DB.Raw(query, request.Email).Scan(&staff).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	if !utils.CheckPasswordHash(request.Password, staff.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token, err := utils.GenerateToken(staff.ID, utils.Config.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
