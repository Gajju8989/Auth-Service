package service

import "time"

const (
	userCreatedMsg    = "User created successfully"
	loginSuccessMsg   = "Login successful"
	userNotFoundMsg   = "User not found"
	jwtExpiryTime     = time.Hour * 1
	refreshExpiryTime = time.Hour * 24 * 7
)

const (
	emailConflictErrMessage = "User with this email already exists"
)
