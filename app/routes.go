package routes

import (
	middlewares "hastane-takip/internal/middleware"
	clinicroutes "hastane-takip/internal/routes/clinic-routes"
	hospitalroutes "hastane-takip/internal/routes/hospital-routes"
	staffroutes "hastane-takip/internal/routes/staff-routes"
	userroutes "hastane-takip/internal/routes/user-routes"

	// Diğer route importları
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func SetupRoutes(app *fiber.App, container *dig.Container) error {
	app.Use(middlewares.AuthMiddleware())
	container.Invoke(func(
		updateClinicRoute *clinicroutes.UpdateClinicRoute,
		deleteClinicRoute *clinicroutes.DeleteClinicRoute,
		getClinicsRoute *clinicroutes.GetClinicsRoute,
		addNewClinicRoute *clinicroutes.AddNewClinicRoute,

	) {
		app.Post("/update-clinic/:id", updateClinicRoute.Handler)
		app.Delete("/delete-clinic/:id", deleteClinicRoute.Handler)
		app.Get("/clinics", getClinicsRoute.Handler)
		app.Post("/add-clinic", addNewClinicRoute.Handler)
	})
	container.Invoke(func(
		staffListRoute *staffroutes.StaffListRoute,
		addNewStaffRoute *staffroutes.AddNewStaffRoute,
		updateStaffRoute *staffroutes.UpdateStaffRoute,
		deleteStaffRoute *staffroutes.DeleteStaffRoute,
	) {
		app.Get("/staffs", staffListRoute.Handler)
		app.Post("/update-staff/:id", updateStaffRoute.Handler)
		app.Delete("/delete-staff/:id", deleteStaffRoute.Handler)
		app.Post("/add-staff", addNewStaffRoute.Handler)
	})

	container.Invoke(func(
		registerRoute *userroutes.RegisterRoute,
		loginRoute *userroutes.LoginRoute,
	) {
		app.Post("/register", registerRoute.Handler)
		app.Post("/login", loginRoute.Handler)
	})
	container.Invoke(func(
		requestResetCodeRoute *userroutes.RequestResetCodeRoute,
		resetPasswordRoute *userroutes.PasswordResetRoute,
	) {
		app.Post("/request-reset-code", requestResetCodeRoute.RequestResetCode)
		app.Post("/reset-password", resetPasswordRoute.Handler)
	})
	container.Invoke(func(
		jobCategoriesRoute *staffroutes.JobCategoriesRoute,
	) {
		app.Get("/job-categories", jobCategoriesRoute.Handler)
	})
	container.Invoke(func(
		getProvincesRoute *hospitalroutes.GetProvincesRoute,
	) {
		app.Get("/provinces", getProvincesRoute.Handler)
	})

	return nil
}
