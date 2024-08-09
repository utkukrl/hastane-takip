package clinicroutes

import (
	trait "hastane-takip/internal/trait"
	"hastane-takip/internal/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type GetClinicsRoute struct {
	dig.In
	ClinicHandler trait.IClinicHandler
	StaffHandler  trait.IStaffHandler
}

func (r *GetClinicsRoute) Handler(c *fiber.Ctx) error {
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
	clinics, err := r.ClinicHandler.GetClinicsByHospitalID(hospitalID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch clinics",
		})
	}
	return c.JSON(clinics)
}
