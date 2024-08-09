package staffroutes

import (
	trait "hastane-takip/internal/trait"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type DeleteStaffRoute struct {
	dig.In
	StaffHandler trait.IStaffHandler
}

func (r *DeleteStaffRoute) Handler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	err = r.StaffHandler.DeleteStaff(uint(id))
	if err != nil {
		if err.Error() == "staff record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Staff record not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete staff record",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Staff record deleted successfully",
	})
}
