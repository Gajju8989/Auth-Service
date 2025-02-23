package repo

import (
	"context"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/token"
	"time"
)

func (i *impl) CreateAccessToken(ctx context.Context, tokenData *token.AccessToken) error {
	return i.db.
		WithContext(ctx).
		Create(tokenData).
		Error
}

func (i *impl) RevokeAccessTokenByUserID(ctx context.Context, userID string) error {
	return i.db.WithContext(ctx).
		Model(&token.AccessToken{}).
		Where("user_id = ?", userID).
		Update("revoked_at", time.Now()).
		Error
}

func (i *impl) GetAccessTokenByTokenID(ctx context.Context, tokenID string) (*token.AccessToken, error) {
	var accessToken token.AccessToken
	err := i.db.
		WithContext(ctx).
		Model(&token.AccessToken{}).
		Where("id = ?", tokenID).
		First(&accessToken).
		Error

	return &accessToken, err
}
