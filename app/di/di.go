package di

import (
	hospitalroutes "hastane-takip/internal/routes/hospital-routes"
	staffroutes "hastane-takip/internal/routes/staff-routes"
	userroutes "hastane-takip/internal/routes/user-routes"
	trait "hastane-takip/internal/trait"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(trait.NewClinicHandler)
	container.Provide(trait.NewHospitalHandler)
	container.Provide(trait.NewStaffHandler)
	container.Provide(userroutes.NewRegisterRoute)
	container.Provide(userroutes.NewLoginRoute)
	container.Provide(trait.NewPasswordResetHandler)
	container.Provide(staffroutes.NewJobCategoriesRoute)
	container.Provide(hospitalroutes.NewGetProvincesRoute)

	return container
}
