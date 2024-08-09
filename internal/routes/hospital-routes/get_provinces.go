package hospitalroutes

import (
	"hastane-takip/internal/cache"

	"github.com/gofiber/fiber/v2"
)

type GetProvincesRoute struct{}

func NewGetProvincesRoute() *GetProvincesRoute {
	return &GetProvincesRoute{}
}

func (r *GetProvincesRoute) Handler(c *fiber.Ctx) error {
	provinces, err := cache.GetProvincesCache()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve job categories",
		})
	}
	return c.JSON(provinces)
}
