package profile

import "main/internal/domain/file"

type ProfileService struct {
	profileRepo ProfileRepository
	fileService *file.FileService
}

func NewProfileService(
	profileRepo ProfileRepository,
	//fileService *file.FileService,
) *ProfileService {
	return &ProfileService{
		profileRepo: profileRepo,
		//fileService: fileService,
	}
}

func (s *ProfileService) CreateProfile(userId string, username string, email string) error {
	profile := &Profile{
		UserId:   userId,
		Username: username,
		Email:    email,
	}

	if err := s.profileRepo.Create(profile); err != nil {
		return err
	}

	return nil
}

func (s *ProfileService) GetOwnProfile(userId string) (*ProfileView, error) {
	profile, err := s.profileRepo.GetByUserID(userId)
	if err != nil {
		return nil, err
	}

	return s.profileToFullView(profile), nil
}

func (s *ProfileService) GetUserProfile(requestingUserId string, targetUsername string) (*ProfileView, error) {
	profile, err := s.profileRepo.GetByUsername(targetUsername)
	if err != nil {
		return nil, err
	}

	if requestingUserId == "" {
		return s.profileToPublicView(profile), nil
	}

	if requestingUserId == profile.UserId {
		return s.profileToFullView(profile), nil
	}

	return s.profileToUserView(profile), nil
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

func (s *ProfileService) profileToFullView(profile *Profile) *ProfileView {
	return &ProfileView{
		ID:          profile.ID,
		UserId:      profile.UserId,
		Username:    profile.Username,
		Email:       profile.Email,
		DisplayName: profile.DisplayName,
		AvatarUrl:   profile.AvatarUrl,
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
		AvatarUrl:   profile.AvatarUrl,
		IsDeleted:   profile.IsDeleted,
	}
}

func (s *ProfileService) profileToPublicView(profile *Profile) *ProfileView {
	return &ProfileView{
		ID:          profile.ID,
		UserId:      profile.UserId,
		Username:    profile.Username,
		DisplayName: profile.DisplayName,
		AvatarUrl:   profile.AvatarUrl,
		IsDeleted:   profile.IsDeleted,
	}
}
