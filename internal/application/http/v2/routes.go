package http

import (
	health_http "main/internal/application/http/v2/health"
	profile_http "main/internal/application/http/v2/profile"
	swagger_http "main/internal/application/http/v2/swagger"

	"go.uber.org/fx"
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	healthRoutes *health_http.HealthRoutes,
	swaggerRoutes *swagger_http.SwaggerRoutes,
	profileRoutes *profile_http.ProfileRoutes,
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
