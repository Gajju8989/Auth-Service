package service

import "time"

const (
	userCreatedMsg           = "User created successfully"
	loginSuccessMsg          = "Login successful"
	userNotFoundMsg          = "User not found"
	tokenRefreshedSuccessMsg = "Token refreshed successfully"
	jwtExpiryTime            = time.Hour * 1
	refreshExpiryTime        = time.Hour * 24 * 7
)

const (
	emailConflictErrMessage       = "User with this email already exists"
	invalidTokenErrMessage        = "Invalid token"
	accessTokenIsInvalidOrRevoked = "access token is invalid or revoked"
	refreshTokenIsAlreadyRevoked  = "refresh token is already revoked"
)
