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
				Message:    invalidTokenErrMessage,
			}
		}

		return err
	}

	if tokenData.RevokedAt != nil {
		return fmt.Errorf(refreshTokenIsAlreadyRevoked)
	}

	err = s.repo.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := s.repo.RevokeAccessTokenByUserID(txCtx, tokenData.UserID); err != nil {
			return err
		}

		if err := s.repo.RevokeRefreshTokenByUserID(txCtx, tokenData.UserID); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
