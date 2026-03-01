package profile_infrastructure

const (
	createProfileQuery = `
		INSERT INTO profiles (id, user_id, username, email, display_name, avatar_url, bio, is_active, is_deleted, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`

	getProfileByIDQuery = `
		SELECT id, user_id, username, email, display_name, avatar_url, bio, is_active, is_deleted, created_at, updated_at
		FROM profiles 
		WHERE id = $1 AND is_deleted = false
	`

	getProfileByUserIDQuery = `
		SELECT id, user_id, username, email, display_name, avatar_url, bio, is_active, is_deleted, created_at, updated_at
		FROM profiles 
		WHERE user_id = $1 AND is_deleted = false
	`

	getProfileByUsernameQuery = `
		SELECT id, user_id, username, email, display_name, avatar_url, bio, is_active, is_deleted, created_at, updated_at
		FROM profiles 
		WHERE username = $1 AND is_deleted = false
	`

	updateProfileQuery = `
		UPDATE profiles 
		SET username = $2, email = $3, display_name = $4, avatar_url = $5, bio = $6
		    is_active = $7, is_deleted = $8, updated_at = $9
		WHERE id = $1 AND is_deleted = false
	`

	deleteProfileQuery = `
		UPDATE profiles 
		SET is_deleted = true, updated_at = $2 
		WHERE id = $1 AND is_deleted = false
	`
)
