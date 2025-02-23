package model

type UserProfileResponse struct {
	ProfileExist bool   `json:"profile_exist"`
	UserEmail    string `json:"email,omitempty"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RevokeTokenRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
}
