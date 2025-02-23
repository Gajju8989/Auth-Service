package repo

import (
	"context"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/refreshtoken"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/token"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/user"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(ctx context.Context, userData *user.User) error
	CreateAccessToken(ctx context.Context, tokenData *token.AccessToken) error
	CreateRefreshToken(ctx context.Context, refreshTokenData *refreshtoken.RefreshToken) error
	GetUserByUserID(ctx context.Context, userID string) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	GetRefreshTokenByToken(ctx context.Context, token string) (*refreshtoken.RefreshToken, error)
	//GetAccessTokenByToken(ctx context.Context, token string) (*token.AccessToken, error)
	WithTransaction(ctx context.Context, fn func(txCtx context.Context) error) error
	UpdateRefreshToken(ctx context.Context, refreshTokenData *refreshtoken.RefreshToken) error
	RevokeAccessTokenByUserID(ctx context.Context, userID string) error
}

type impl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &impl{db: db}
}

func (i *impl) WithTransaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	return i.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, "tx", tx)
		return fn(txCtx)
	})
}
