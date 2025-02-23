package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/refreshtoken"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/token"
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

func (s *impl) hashPassword(input string) (string, error) {
	hashedInput, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedInput), nil
}

func (s *impl) createTokens(ctx context.Context, accessTokenUUID, refreshTokenUUID, userID string) error {
	return s.repo.WithTransaction(ctx, func(txCtx context.Context) error {
		err := s.repo.CreateAccessToken(txCtx, &token.AccessToken{
			ID:        accessTokenUUID,
			UserID:    userID,
			ExpiresAt: time.Now().Add(jwtExpiryTime),
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}

		err = s.repo.CreateRefreshToken(txCtx, &refreshtoken.RefreshToken{
			ID:        refreshTokenUUID,
			UserID:    userID,
			ExpiresAt: time.Now().Add(refreshExpiryTime),
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
