package profile_infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	domain "main/internal/domain/profile"
	"main/pkg"
)

type profileRepository struct {
	db pkg.PostgresDB
}

func NewProfileRepository(db pkg.PostgresDB) domain.ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) Create(profile *domain.Profile) error {
	profile.ID = generateProfileID()
	profile.CreatedAt = time.Now().UTC()
	profile.UpdatedAt = time.Now().UTC()

	var returnedID string
	err := r.db.QueryRow(
		createProfileQuery,
		profile.ID,
		profile.UserId,
		profile.Username,
		profile.Email,
		profile.DisplayName,
		profile.AvatarUrl,
		profile.Bio,
		profile.IsActive,
		profile.IsDeleted,
		profile.CreatedAt,
		profile.UpdatedAt,
	).Scan(&returnedID)

	if err != nil {
		return err
	}

	return nil
}

func (r *profileRepository) GetById(profileId string) (*domain.Profile, error) {
	profile := &domain.Profile{}
	err := r.db.QueryRow(getProfileByIDQuery, profileId).Scan(
		&profile.ID,
		&profile.UserId,
		&profile.Username,
		&profile.Email,
		&profile.DisplayName,
		&profile.AvatarUrl,
		&profile.Bio,
		&profile.IsActive,
		&profile.IsDeleted,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrProfileNotFound
		}
		return nil, err
	}

	return profile, nil
}

func (r *profileRepository) GetByUserID(userID string) (*domain.Profile, error) {
	profile := &domain.Profile{}
	err := r.db.QueryRow(getProfileByUserIDQuery, userID).Scan(
		&profile.ID,
		&profile.UserId,
		&profile.Username,
		&profile.Email,
		&profile.DisplayName,
		&profile.AvatarUrl,
		&profile.Bio,
		&profile.IsActive,
		&profile.IsDeleted,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrProfileNotFound
		}
		return nil, err
	}

	return profile, nil
}

func (r *profileRepository) GetByUsername(username string) (*domain.Profile, error) {
	if username == "" {
		return nil, domain.ErrProfileNotFound
	}

	profile := &domain.Profile{}
	err := r.db.QueryRow(getProfileByUsernameQuery, username).Scan(
		&profile.ID,
		&profile.UserId,
		&profile.Username,
		&profile.Email,
		&profile.DisplayName,
		&profile.AvatarUrl,
		&profile.Bio,
		&profile.IsActive,
		&profile.IsDeleted,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrProfileNotFound
		}
		return nil, err
	}

	return profile, nil
}

func (r *profileRepository) Update(profile *domain.Profile) error {
	profile.UpdatedAt = time.Now().UTC()

	result, err := r.db.Exec(
		updateProfileQuery,
		profile.ID,
		profile.Username,
		profile.Email,
		profile.DisplayName,
		profile.AvatarUrl,
		profile.Bio,
		profile.IsActive,
		profile.IsDeleted,
		profile.UpdatedAt,
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.ErrProfileNotFound
	}

	return nil
}

func (r *profileRepository) Delete(profileId string) error {
	result, err := r.db.Exec(deleteProfileQuery, profileId, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.ErrProfileNotFound
	}

	return nil
}

func (r *profileRepository) DeleteByUserId(userId string) error {
	result, err := r.db.Exec(deleteProfileByUserIdQuery, userId, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.ErrProfileNotFound
	}

	return nil
}

func generateProfileID() string {
	return fmt.Sprintf("prof_%d", time.Now().UnixNano())
}
