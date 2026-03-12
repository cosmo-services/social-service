package bootstrap

import (
	"main/internal/application/http/v2"
	"main/internal/application/jobs"
	"main/internal/config"
	"main/internal/domain"
	"main/pkg"

	auth_infrastructure "main/internal/infrastructure/auth"
	profile_infrastructure "main/internal/infrastructure/profile"

	profile_domain "main/internal/domain/profile"

	auth_http "main/internal/application/http/v2/auth"
	health_http "main/internal/application/http/v2/health"
	profile_http "main/internal/application/http/v2/profile"
	swagger_http "main/internal/application/http/v2/swagger"

	nats "main/internal/application/nats"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	pkg.Module,
	domain.Module,

	profile_domain.Module,

	auth_infrastructure.Module,
	profile_infrastructure.Module,

	http.Module,
	jobs.Module,
	nats.Module,
	health_http.Module,
	swagger_http.Module,
	profile_http.Module,
	auth_http.Module,
)
