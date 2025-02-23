package service

import (
	"context"
	"errors"
	"github/com/Gajju8989/Auth_Service/internal/service/model"
	"gorm.io/gorm"
	"net/http"
)

func (s *impl) GetProfiles(ctx context.Context, userID string) (*model.UserProfileResponse, error) {
	token, err := s.repo.GetAccessTokenByTokenID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &GenericResponse{
				StatusCode: http.StatusForbidden,
				Message:    "invalid token",
			}
		}

		return nil, err
	}

	// If the token is not found or is revoked, return an error
	if token == nil || token.RevokedAt != nil {
		return nil, &GenericResponse{
			StatusCode: http.StatusForbidden,
			Message:    "access token is invalid or revoked",
		}
	}

	userData, err := s.repo.GetUserByUserID(ctx, token.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &model.UserProfileResponse{
				ProfileExist: false,
			}, nil
		}

		return nil, err
	}

	return &model.UserProfileResponse{
		ProfileExist: true,
		UserEmail:    userData.Email,
	}, nil
}
