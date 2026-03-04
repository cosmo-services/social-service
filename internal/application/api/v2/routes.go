package api

import (
	health_api "main/internal/application/api/v2/health"
	profile_api "main/internal/application/api/v2/profile"
	swagger_api "main/internal/application/api/v2/swagger"

	"go.uber.org/fx"
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	healthRoutes *health_api.HealthRoutes,
	swaggerRoutes *swagger_api.SwaggerRoutes,
	profileRoutes *profile_api.ProfileRoutes,
) Routes {
	return Routes{
		healthRoutes,
		swaggerRoutes,
		profileRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

var Module = fx.Options(
	fx.Provide(NewRoutes),
)
