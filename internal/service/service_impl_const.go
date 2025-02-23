package service

import "time"

const (
	userCreatedMsg    = "User created successfully"
	loginSuccessMsg   = "Login successful"
	jwtExpiryTime     = time.Hour * 1
	refreshExpiryTime = time.Hour * 24 * 7
)
