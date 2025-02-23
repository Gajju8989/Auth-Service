package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *impl) generateAccessToken(userID string) (string, error) {
	accessTokenClaims := &jwt.StandardClaims{
		Subject:   userID,
		Id:        uuid.New().String(),
		ExpiresAt: time.Now().Add(jwtExpiryTime).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func (s *impl) generateRefreshToken(userID string) (string, error) {
	refreshTokenClaims := &jwt.StandardClaims{
		Subject:   userID,
		Id:        uuid.New().String(),
		ExpiresAt: time.Now().Add(refreshExpiryTime).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

func (s *impl) hashInput(input string) (string, error) {
	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte(input))
	hashedInput := sha256Hasher.Sum(nil)
	hashedInputStr := hex.EncodeToString(hashedInput)

	bcryptHashedInput, err := bcrypt.GenerateFromPassword([]byte(hashedInputStr), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bcryptHashedInput), nil
}

func (s *impl) hashPassword(input string) (string, error) {
	hashedInput, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedInput), nil
}
