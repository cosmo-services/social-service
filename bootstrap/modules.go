package bootstrap

import (
	"main/internal/application/api/v2"
	"main/internal/application/jobs"
	"main/internal/config"
	"main/internal/domain"
	"main/pkg"

	auth_infrastructure "main/internal/infrastructure/auth"
	profile_infrastructure "main/internal/infrastructure/profile"
	test_infrastructure "main/internal/infrastructure/test"

	//file_domain "main/internal/domain/file"
	profile_domain "main/internal/domain/profile"

	auth_api "main/internal/application/api/v2/auth"
	health_api "main/internal/application/api/v2/health"
	swagger_api "main/internal/application/api/v2/swagger"
	test_api "main/internal/application/api/v2/test"

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
	test_infrastructure.Module,
	profile_infrastructure.Module,

	api.Module,
	jobs.Module,
	nats.Module,
	health_api.Module,
	swagger_api.Module,
	test_api.Module,
	auth_api.Module,
)
