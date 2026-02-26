package profile_api

import (
	"main/internal/application/api/v2/auth"
	"main/pkg"
)

type ProfileRoutes struct {
	handler           pkg.RequestHandler
	profileController *ProfileController
	authMiddleware    *auth.AuthMiddleware
}

func NewProfileRoutes(
	profileController *ProfileController,
	handler pkg.RequestHandler,
	authMiddleware *auth.AuthMiddleware,
) *ProfileRoutes {
	return &ProfileRoutes{
		profileController: profileController,
		handler:           handler,
		authMiddleware:    authMiddleware,
	}
}

func (r *ProfileRoutes) Setup() {
	api := r.handler.Gin.Group("/api/v2/social/profile")

	protected := api.Group("/")
	protected.Use(r.authMiddleware.RequireAuth())
	{
		protected.GET("/me", r.profileController.GetOwnProfile)
		protected.GET("/:username", r.profileController.GetUserProfile)
	}
}
