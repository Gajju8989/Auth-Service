package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github/com/Gajju8989/Auth_Service/internal/service/model"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func (s *impl) RefreshToken(ctx context.Context, refreshTokenID string) (*model.UserAuthResponse, error) {
	refreshTokenData, err := s.repo.GetRefreshTokenByID(ctx, refreshTokenID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &GenericResponse{
				StatusCode: http.StatusForbidden,
				Message:    invalidTokenErrMessage,
			}
		}

		return nil, err
	}

	if refreshTokenData.RevokedAt != nil || refreshTokenData.ExpiresAt.Before(time.Now()) {
		return nil, &GenericResponse{
			StatusCode: http.StatusForbidden,
			Message:    accessTokenIsInvalidOrRevoked,
		}
	}

	accessTokenUUID := uuid.New().String()
	accessToken, err := s.generateAccessToken(accessTokenUUID)
	if err != nil {
		return nil, err
	}

	refreshTokenUUID := uuid.New().String()
	refreshToken, err := s.generateRefreshToken(refreshTokenUUID)
	if err != nil {
		return nil, err
	}

	err = s.createTokens(ctx, accessTokenUUID, refreshTokenUUID, refreshTokenData.UserID)
	if err != nil {
		return nil, err
	}

	return &model.UserAuthResponse{
		Message: tokenRefreshedSuccessMsg,
		Token: &model.Tokens{
			AccessToken:        accessToken,
			AccessTokenExpiry:  time.Now().Add(jwtExpiryTime).Unix(),
			RefreshToken:       refreshToken,
			RefreshTokenExpiry: time.Now().Add(refreshExpiryTime).Unix(),
		},
	}, nil
}
