package repo

import (
	"context"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/refreshtoken"
	"time"
)

func (i *impl) CreateRefreshToken(ctx context.Context, refreshTokenData *refreshtoken.RefreshToken) error {
	return i.db.
		WithContext(ctx).
		Create(refreshTokenData).
		Error
}

func (i *impl) GetRefreshTokenByID(ctx context.Context, refreshTokenID string) (refreshtoken.RefreshToken, error) {
	var refreshToken refreshtoken.RefreshToken
	err := i.db.
		WithContext(ctx).
		Where("id = ?", refreshTokenID).
		First(&refreshToken).
		Error

	return refreshToken, err
}

func (i *impl) RevokeRefreshTokenByUserID(ctx context.Context, userID string) error {
	return i.db.WithContext(ctx).
		Model(&refreshtoken.RefreshToken{}).
		Where("user_id = ?", userID).
		Update("revoked_at", time.Now()).
		Error
}
