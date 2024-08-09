package hospitalroutes

import (
	trait "hastane-takip/internal/trait"
	"hastane-takip/internal/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type GetHospitalsRoute struct {
	dig.In
	HospitalHandler trait.IHospitalHandler
}

func (r *GetHospitalsRoute) Handler(c *fiber.Ctx) error {
	userID, err := utils.GetStaffID(c)
	hospitals, err := r.HospitalHandler.GetHospitalByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch hospita",
		})
	}
	return c.JSON(hospitals)
}
