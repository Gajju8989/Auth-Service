package repo

import (
	"context"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/refreshtoken"
)

func (i *impl) CreateRefreshToken(ctx context.Context, refreshTokenData *refreshtoken.RefreshToken) error {
	return i.db.
		WithContext(ctx).
		Create(refreshTokenData).
		Error
}

func (i *impl) GetRefreshTokenByToken(ctx context.Context, token string) (*refreshtoken.RefreshToken, error) {
	var refreshToken refreshtoken.RefreshToken
	err := i.db.
		WithContext(ctx).
		Where("refresh_token_hash = ?", token).
		First(&refreshToken).
		Error

	return &refreshToken, err
}

func (i *impl) UpdateRefreshToken(ctx context.Context, refreshTokenData *refreshtoken.RefreshToken) error {
	return i.db.WithContext(ctx).
		Model(&refreshtoken.RefreshToken{}).
		Where("id = ?", refreshTokenData.ID).
		Updates(refreshTokenData).
		Error
}
