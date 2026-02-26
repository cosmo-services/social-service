package bootstrap

import (
	"main/internal/application/api/v2"
	"main/internal/application/jobs"
	"main/internal/config"
	"main/internal/domain"
	"main/pkg"

	auth_infrastructure "main/internal/infrastructure/auth"
	profile_infrastructure "main/internal/infrastructure/profile"

	//file_domain "main/internal/domain/file"
	profile_domain "main/internal/domain/profile"

	auth_api "main/internal/application/api/v2/auth"
	health_api "main/internal/application/api/v2/health"
	profile_api "main/internal/application/api/v2/profile"
	swagger_api "main/internal/application/api/v2/swagger"

	nats "main/internal/application/nats"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	pkg.Module,
	domain.Module,

	profile_domain.Module,
	//file_domain.Module,

	auth_infrastructure.Module,
	profile_infrastructure.Module,

	api.Module,
	jobs.Module,
	nats.Module,
	health_api.Module,
	swagger_api.Module,
	profile_api.Module,
	auth_api.Module,
)
