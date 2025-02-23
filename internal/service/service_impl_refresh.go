package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/refreshtoken"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/token"
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
				Message:    "Invalid token",
			}
		}

		return nil, err
	}

	if refreshTokenData.RevokedAt != nil || refreshTokenData.ExpiresAt.Before(time.Now()) {
		return nil, &GenericResponse{
			StatusCode: http.StatusForbidden,
			Message:    "Token is revoked or expired",
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

	hashedAccessToken, err := s.hashInput(accessToken)
	if err != nil {
		return nil, err
	}

	hashedRefreshToken, err := s.hashInput(refreshToken)
	if err != nil {
		return nil, err
	}

	err = s.repo.WithTransaction(ctx, func(txCtx context.Context) error {
		err = s.repo.CreateAccessToken(txCtx, &token.AccessToken{
			ID:        accessTokenUUID,
			UserID:    refreshTokenData.UserID,
			TokenHash: hashedAccessToken,
			ExpiresAt: time.Now().Add(jwtExpiryTime),
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}

		err = s.repo.CreateRefreshToken(txCtx, &refreshtoken.RefreshToken{
			ID:               refreshTokenUUID,
			UserID:           refreshTokenData.UserID,
			RefreshTokenHash: hashedRefreshToken,
			ExpiresAt:        time.Now().Add(refreshExpiryTime),
			CreatedAt:        time.Now(),
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model.UserAuthResponse{
		Message: "Token refreshed successfully",
		Token: &model.Tokens{
			AccessToken:        accessToken,
			AccessTokenExpiry:  time.Now().Add(jwtExpiryTime).Unix(),
			RefreshToken:       refreshToken,
			RefreshTokenExpiry: time.Now().Add(refreshExpiryTime).Unix(),
		},
	}, nil
}
