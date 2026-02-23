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
