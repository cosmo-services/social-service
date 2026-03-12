package profile_api

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewProfileController),
	fx.Provide(NewProfileRoutes),
)
