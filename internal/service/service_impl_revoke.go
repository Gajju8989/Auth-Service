package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

func (s *impl) RevokeToken(ctx context.Context, userID string) error {
	tokenData, err := s.repo.GetAccessTokenByTokenID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &GenericResponse{
				StatusCode: http.StatusForbidden,
				Message:    "Invalid token",
			}
		}

		return err
	}

	if tokenData.RevokedAt != nil {
		return fmt.Errorf("refresh token is already revoked")
	}

	err = s.repo.WithTransaction(ctx, func(txCtx context.Context) error {
		err = s.repo.RevokeAccessTokenByUserID(txCtx, tokenData.UserID)
		if err != nil {
			return err
		}

		err = s.repo.RevokeRefreshTokenByUserID(txCtx, tokenData.UserID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
