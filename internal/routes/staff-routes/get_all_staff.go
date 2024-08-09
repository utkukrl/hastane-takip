package staffroutes

import (
	trait "hastane-takip/internal/trait"
	"hastane-takip/internal/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type StaffListRoute struct {
	dig.In
	StaffHandler trait.IStaffHandler
}

func (r *StaffListRoute) Handler(c *fiber.Ctx) error {
	userID, err := utils.GetStaffID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	hospitalID, err := r.StaffHandler.GetHospitalIDByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve hospital ID",
		})
	}

	staffList, err := r.StaffHandler.GetStaffByHospitalID(hospitalID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve staff list",
		})
	}
	return c.Status(fiber.StatusOK).JSON(staffList)
}
