package profile

import (
	"main/internal/domain"
	"time"
)

type ProfileService struct {
	profileRepo ProfileRepository
	eventBus    *domain.EventBus
}

func NewProfileService(
	profileRepo ProfileRepository,
	eventBus *domain.EventBus,
) *ProfileService {
	return &ProfileService{
		profileRepo: profileRepo,
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

func (s *ProfileService) GetProfileById(requestingUserId string, targetUserId string) (*ProfileView, error) {
	profile, err := s.profileRepo.GetByUserID(targetUserId)
	if err != nil {
		return nil, err
	}

	return s.profileToView(requestingUserId, profile)
}

func (s *ProfileService) GetProfileByUsername(requestingUserId string, targetUsername string) (*ProfileView, error) {
	profile, err := s.profileRepo.GetByUsername(targetUsername)
	if err != nil {
		return nil, err
	}

	return s.profileToView(requestingUserId, profile)
}

func (s *ProfileService) Activate(userId string) error {
	profile, err := s.profileRepo.GetByUserID(userId)
	if err != nil {
		return err
	}

	profile.IsActive = true
	if err := s.profileRepo.Update(profile); err != nil {
		return err
	}

	return nil
}

func (s *ProfileService) UpdateEmail(userId string, newEmail string) error {
	profile, err := s.profileRepo.GetByUserID(userId)
	if err != nil {
		return err
	}

	profile.Email = newEmail
	if err := s.profileRepo.Update(profile); err != nil {
		return err
	}

	return nil
}

func (s *ProfileService) UpdateUsername(userId string, newUsername string) error {
	profile, err := s.profileRepo.GetByUserID(userId)
	if err != nil {
		return err
	}

	profile.Username = newUsername
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
