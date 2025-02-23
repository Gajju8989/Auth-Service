package model

type UserAuthRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserAuthResponse struct {
	Message string  `json:"message"`
	Token   *Tokens `json:"token,omitempty"`
}

type Tokens struct {
	AccessToken        string `json:"access_token"`
	AccessTokenExpiry  int64  `json:"access_token_expiry"`
	RefreshToken       string `json:"refresh_token"`
	RefreshTokenExpiry int64  `json:"refresh_token_expiry"`
}
