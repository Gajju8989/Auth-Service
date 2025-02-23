package service

import (
	"context"
	"github/com/Gajju8989/Auth_Service/internal/repo"
	"github/com/Gajju8989/Auth_Service/internal/service/model"
)

type AuthService interface {
	CreateUser(ctx context.Context, req *model.UserAuthRequest) (*model.UserAuthResponse, error)
	Login(ctx context.Context, req *model.UserAuthRequest) (*model.UserAuthResponse, error)
	GetProfiles(ctx context.Context, userID string) (*model.UserProfileResponse, error)
	RefreshToken(ctx context.Context, req *model.RefreshTokenRequest) (*model.UserAuthResponse, error)
	RevokeToken(ctx context.Context, req *model.RevokeTokenRequest) error
}

type impl struct {
	repo repo.Repository
}

func NewAuthService(repo repo.Repository) AuthService {
	return &impl{repo: repo}
}
