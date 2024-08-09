package staffroutes

import (
	"hastane-takip/internal/models"
	trait "hastane-takip/internal/trait"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type AddNewStaffRoute struct {
	dig.In
	StaffHandler trait.IStaffHandler
}

func (r *AddNewStaffRoute) Handler(c *fiber.Ctx) error {
	var addedStaff models.Staff

	if err := c.BodyParser(&addedStaff); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := r.StaffHandler.CreateStaff(&addedStaff)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not add staff",
		})
	}
	return c.JSON(fiber.Map{})

}
