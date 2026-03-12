package grpc_v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/cosmo-services/grpc-contracts/gen/api/v1"

	profile_domain "main/internal/domain/profile"
)

type SocialHandler struct {
	pb.UnimplementedSocialServiceServer

	profileService *profile_domain.ProfileService
}

func NewSocialHandler(
	profileService *profile_domain.ProfileService,
) *SocialHandler {
	return &SocialHandler{
		profileService: profileService,
	}
}

func (h *SocialHandler) GetUserProfile(
	ctx context.Context,
	req *pb.GetUserProfileRequest,
) (*pb.GetUserProfileResponse, error) {
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	p, err := h.profileService.GetProfile(profile_domain.ProfileSearchOptions{UserID: req.UserId})
	if err != nil {
		if errors.Is(err, profile_domain.ErrProfileNotFound) {
			return nil, status.Error(codes.NotFound, "profile not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetUserProfileResponse{
		Profile: &pb.UserProfile{
			Id:          p.ID,
			Username:    p.Username,
			AvatarUrl:   p.AvatarUrl,
			DisplayName: p.DisplayName,
		},
	}, nil
}
