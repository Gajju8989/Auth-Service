package service

import (
	"context"
	"fmt"
	
	"github/com/Gajju8989/Auth_Service/internal/config/jwtkey"
	"github/com/Gajju8989/Auth_Service/internal/repo"
	"github/com/Gajju8989/Auth_Service/internal/service/model"
)

type AuthService interface {
	CreateUser(ctx context.Context, req *model.UserAuthRequest) (*model.UserAuthResponse, error)
	Login(ctx context.Context, req *model.UserAuthRequest) (*model.UserAuthResponse, error)
	GetProfiles(ctx context.Context, userID string) (*model.UserProfileResponse, error)
	RefreshToken(ctx context.Context, userID string) (*model.UserAuthResponse, error)
	RevokeToken(ctx context.Context, userID string) error
}

type impl struct {
	repo   repo.Repository
	jwtKey []byte
}

func NewAuthService(repo repo.Repository) (AuthService, error) {
	jwtKey, err := jwtkey.GetJWTKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get JWT key: %w", err)
	}

	return &impl{repo: repo, jwtKey: jwtKey}, nil
}
