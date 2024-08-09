package staffroutes

import (
	"hastane-takip/internal/cache"

	"github.com/gofiber/fiber/v2"
)

type JobCategoriesRoute struct{}

func NewJobCategoriesRoute() *JobCategoriesRoute {
	return &JobCategoriesRoute{}
}

func (r *JobCategoriesRoute) Handler(c *fiber.Ctx) error {
	jobCategories, err := cache.GetJobCategoriesCache()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve job categories",
		})
	}
	return c.JSON(jobCategories)
}
