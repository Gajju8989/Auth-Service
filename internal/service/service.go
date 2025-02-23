package service

import (
	"context"
	"github.com/joho/godotenv"
	"github/com/Gajju8989/Auth_Service/internal/repo"
	"github/com/Gajju8989/Auth_Service/internal/service/model"
	"os"
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

func NewAuthService(repo repo.Repository) AuthService {
	err := godotenv.Load("local.env")
	if err != nil {
		panic("Error loading .env file")
	}

	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	return &impl{repo: repo, jwtKey: jwtKey}
}
