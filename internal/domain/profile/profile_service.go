package profile

import (
	"errors"
	"main/internal/domain"

	"main/internal/domain/auth"

	"time"
)

type ProfileService struct {
	profileRepo ProfileRepository
	authClient  auth.AuthClient
	eventBus    *domain.EventBus
}

func NewProfileService(
	profileRepo ProfileRepository,
	authClient auth.AuthClient,
	eventBus *domain.EventBus,
) *ProfileService {
	return &ProfileService{
		profileRepo: profileRepo,
		authClient:  authClient,
		eventBus:    eventBus,
	}
}

func (s *ProfileService) CreateProfile(userId string, username string, email string) error {
	profile := &Profile{
		UserId:      userId,
		Username:    username,
		Email:       email,
		DisplayName: username,
	}

	if err := s.profileRepo.Create(profile); err != nil {
		return err
	}

	return nil
}

func (s *ProfileService) GetProfile(opts ProfileSearchOptions) (*Profile, error) {
	var profile *Profile
	var err error

	if opts.UserID != "" {
		profile, err = s.profileRepo.GetByUserID(opts.UserID)
	} else if opts.Username != "" {
		profile, err = s.profileRepo.GetByUsername(opts.Username)
	} else if opts.ProfileID != "" {
		profile, err = s.profileRepo.GetById(opts.ProfileID)
	}

	if err == nil {
		return profile, nil
	}

	if !errors.Is(err, ErrProfileNotFound) {
		return nil, err
	}

	var user *auth.AuthUser
	if opts.UserID != "" {
		user, err = s.authClient.GetUserById(opts.UserID)
	} else if opts.Username != "" {
		user, err = s.authClient.GetUserByUsername(opts.Username)
	}

	if err != nil {
		return nil, err
	}

	profile = MapUserToProfile(user)

	if err := s.profileRepo.Create(profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *ProfileService) GetProfileViewById(requestingUserId string, targetUserId string) (*ProfileView, error) {
	profile, err := s.GetProfile(ProfileSearchOptions{UserID: targetUserId})
	if err != nil {
		return nil, err
	}

	return s.profileToView(requestingUserId, profile)
}

func (s *ProfileService) GetProfileViewByUsername(requestingUserId string, targetUsername string) (*ProfileView, error) {
	profile, err := s.GetProfile(ProfileSearchOptions{Username: targetUsername})
	if err != nil {
		return nil, err
	}

	return s.profileToView(requestingUserId, profile)
}

func (s *ProfileService) UpdateUser(userId string, username string, email string, isActive bool) error {
	profile, err := s.profileRepo.GetByUserID(userId)
	if err != nil {
		return err
	}

	profile.Username = username
	profile.IsActive = isActive
	profile.Email = email

	if err := s.profileRepo.Update(profile); err != nil {
		return err
	}

	return nil
}

func (s *ProfileService) ChangeBio(userId string, newBio string) error {
	profile, err := s.profileRepo.GetByUserID(userId)
	if err != nil {
		return err
	}

	if err := profile.ChangeBio(newBio); err != nil {
		return err
	}

	if err := s.profileRepo.Update(profile); err != nil {
		return err
	}

	return nil
}

func (s *ProfileService) ChangeAvatar(userId string, newAvatar string) error {
	profile, err := s.profileRepo.GetByUserID(userId)
	if err != nil {
		return err
	}

	oldAvatar := profile.AvatarUrl
	if err := profile.ChangeAvatar(newAvatar); err != nil {
		return err
	}

	if err := s.profileRepo.Update(profile); err != nil {
		return err
	}

	s.eventBus.Emit("avatar.changed", ProfileAvatarChangedEvent{
		NewAvatar: newAvatar,
		ChangedAt: time.Now(),
	})

	s.eventBus.Emit("avatar.orphaned", FileOrphanedEvent{
		FilePath:   oldAvatar,
		OrphanedAt: time.Now(),
	})

	return nil
}

func (s *ProfileService) DeleteAvatar(userId string, newAvatar string) error {
	profile, err := s.profileRepo.GetByUserID(userId)
	if err != nil {
		return err
	}

	oldAvatar := profile.AvatarUrl
	if err := profile.ChangeAvatar(newAvatar); err != nil {
		return err
	}

	if err := s.profileRepo.Update(profile); err != nil {
		return err
	}

	s.eventBus.Emit("avatar.changed", ProfileAvatarChangedEvent{
		NewAvatar: newAvatar,
		ChangedAt: time.Now(),
	})

	s.eventBus.Emit("avatar.orphaned", FileOrphanedEvent{
		FilePath:   oldAvatar,
		OrphanedAt: time.Now(),
	})

	return nil
}

func (s *ProfileService) ChangeDisplayName(userId string, newDisplayName string) error {
	profile, err := s.profileRepo.GetByUserID(userId)
	if err != nil {
		return err
	}

	if err := profile.ChangeDisplayName(newDisplayName); err != nil {
		return err
	}

	if err := s.profileRepo.Update(profile); err != nil {
		return err
	}

	return nil
}

func (s *ProfileService) DeleteProfile(userId string) error {
	if err := s.profileRepo.DeleteByUserId(userId); err != nil {
		return err
	}

	return nil
}

func (s *ProfileService) profileToView(requestingUserId string, profile *Profile) (*ProfileView, error) {
	if requestingUserId == "" {
		return s.profileToPublicView(profile), nil
	}

	if requestingUserId == profile.UserId {
		return s.profileToFullView(profile), nil
	}

	return s.profileToUserView(profile), nil
}

func (s *ProfileService) profileToFullView(profile *Profile) *ProfileView {
	return &ProfileView{
		ID:          profile.ID,
		UserId:      profile.UserId,
		Username:    profile.Username,
		Email:       profile.Email,
		DisplayName: profile.DisplayName,
		AvatarUrl:   profile.AvatarUrl,
		Bio:         profile.Bio,
		IsActive:    profile.IsActive,
		IsDeleted:   profile.IsDeleted,
		CreatedAt:   profile.CreatedAt.UTC().String(),
	}
}

func (s *ProfileService) profileToUserView(profile *Profile) *ProfileView {
	return &ProfileView{
		ID:          profile.ID,
		UserId:      profile.UserId,
		Username:    profile.Username,
		DisplayName: profile.DisplayName,
		Bio:         profile.Bio,
		AvatarUrl:   profile.AvatarUrl,
		IsDeleted:   profile.IsDeleted,
		CreatedAt:   profile.CreatedAt.UTC().String(),
	}
}

func (s *ProfileService) profileToPublicView(profile *Profile) *ProfileView {
	return &ProfileView{
		ID:          profile.ID,
		UserId:      profile.UserId,
		Username:    profile.Username,
		DisplayName: profile.DisplayName,
		Bio:         profile.Bio,
		AvatarUrl:   profile.AvatarUrl,
		IsDeleted:   profile.IsDeleted,
		CreatedAt:   profile.CreatedAt.UTC().String(),
	}
}

func (s *ProfileService) requestProfileByUserId(userId string) (*Profile, error) {
	user, err := s.authClient.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	profile := MapUserToProfile(user)
	return profile, nil
}

func (s *ProfileService) requestProfileByUsername(username string) (*Profile, error) {
	user, err := s.authClient.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	profile := MapUserToProfile(user)
	return profile, nil
}
