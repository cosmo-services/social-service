package profile_api

type ChangeBioRequest struct {
	NewBio string `json:"new_bio"`
}

type ChangeDisplayNameRequest struct {
	NewDisplayName string `json:"new_display_name"`
}
