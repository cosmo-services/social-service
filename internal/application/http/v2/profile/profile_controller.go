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
	requestingUserId := ctx.GetString("user_id")
	targetUserId := requestingUserId

	profile, err := controller.profileService.GetProfileViewById(requestingUserId, targetUserId)
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
// @Description Get user's profile data by ID or username.
// @Tags profile
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param user_id query string false "User ID"
// @Param username query string false "Username"
// @Success 200 {object} profile.Profile
// @Failure 400 {object} map[string]string "Bad request - missing both parameters or both provided"
// @Failure 500 {object} map[string]string "Internal server error"
// @Failure 404 {object} map[string]string "Profile not found"
// @Router /profile [get]
func (controller *ProfileController) GetUserProfile(ctx *gin.Context) {
	requestingUserId := ctx.GetString("user_id")
	targetUserId := ctx.Query("user_id")
	targetUsername := ctx.Query("username")

	if targetUserId == "" && targetUsername == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "either user_id or username query parameter is required"})
		return
	}

	if targetUserId != "" && targetUsername != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "provide only one: user_id or username, not both"})
		return
	}

	var profile *profile_domain.ProfileView
	var err error

	if targetUserId != "" {
		profile, err = controller.profileService.GetProfileViewById(requestingUserId, targetUserId)
	} else {
		profile, err = controller.profileService.GetProfileViewByUsername(requestingUserId, targetUsername)
	}

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

// ChangeBio godoc
//
// @Summary Change bio
// @Description Change user's profile bio.
// @Tags profile
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param request body	ChangeBioRequest true "Change bio request"
// @Success 200 {object} map[string]string "Bio changed successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Failure 404 {object} map[string]string "Profile not found"
// @Failure 401 {object} map[string]string "User unauthorized"
// @Router /profile/bio [post]
func (controller *ProfileController) ChangeBio(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	var req *ChangeBioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := controller.profileService.ChangeBio(userId, req.NewBio); err != nil {
		if errors.Is(err, profile_domain.ErrProfileNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, profile_domain.ErrLongBio) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "bio changed successfully",
	})
}

// ChangeDisplayName godoc
//
// @Summary Change display name
// @Description Change user's display name.
// @Tags profile
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param request body	ChangeDisplayNameRequest true "Change display name request"
// @Success 200 {object} map[string]string "Display name changed successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Failure 404 {object} map[string]string "Profile not found"
// @Failure 401 {object} map[string]string "User unauthorized"
// @Router /profile/diaplayname [post]
func (controller *ProfileController) ChangeDisplayName(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	var req *ChangeDisplayNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := controller.profileService.ChangeDisplayName(userId, req.NewDisplayName); err != nil {
		if errors.Is(err, profile_domain.ErrProfileNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, profile_domain.ErrLongDisplayName) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "display name changed successfully",
	})
}
