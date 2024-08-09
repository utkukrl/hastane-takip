package staffroutes

import (
	"hastane-takip/internal/models"
	trait "hastane-takip/internal/trait"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type UpdateStaffRoute struct {
	dig.In
	StaffHandler trait.IStaffHandler
}

func (r *UpdateStaffRoute) Handler(c *fiber.Ctx) error {
	var staffUpdate models.Staff

	if err := c.BodyParser(&staffUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	err = r.StaffHandler.UpdateStaff(uint(id), &staffUpdate)
	if err != nil {
		if err.Error() == "staff record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Staff record not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update staff record",
		})
	}

	return c.Status(fiber.StatusOK).JSON(staffUpdate)

}
