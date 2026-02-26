package profile_api

import (
	"errors"
	profile_domain "main/internal/domain/profile"
	"main/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileService *profile_domain.ProfileService
	logger         pkg.Logger
}

func NewProfileController(
	profileService *profile_domain.ProfileService,
	logger pkg.Logger,
) *ProfileController {
	return &ProfileController{
		profileService: profileService,
		logger:         logger,
	}
}

// GetOwnProfile godoc
//
// @Summary Get current user profile
// @Description Get the authenticated user's profile data.
// @Tags profile
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} profile.Profile
// @Failure 500 {object} map[string]string "Internal server error"
// @Failure 404 {object} map[string]string "Profile not found"
// @Failure 401 {object} map[string]string "User unauthorized"
// @Router /profile/me [get]
func (controller *ProfileController) GetOwnProfile(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	profile, err := controller.profileService.GetOwnProfile(userId)
	if err != nil {
		if errors.Is(err, profile_domain.ErrProfileNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

// GetUserProfile godoc
//
// @Summary Get user profile
// @Description Get user's profile data.
// @Tags profile
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param username path string true "Username"
// @Success 200 {object} profile.Profile
// @Failure 500 {object} map[string]string "Internal server error"
// @Failure 404 {object} map[string]string "Profile not found"
// @Router /profile/{username} [get]
func (controller *ProfileController) GetUserProfile(ctx *gin.Context) {
	requestingUserId := ctx.GetString("user_id")
	targetUsername := ctx.Param("username")

	profile, err := controller.profileService.GetUserProfile(requestingUserId, targetUsername)
	if err != nil {
		if errors.Is(err, profile_domain.ErrProfileNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}
