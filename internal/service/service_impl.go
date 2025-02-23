package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/refreshtoken"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/token"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/user"
	"github/com/Gajju8989/Auth_Service/internal/service/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func (s *impl) CreateUser(ctx context.Context, req *model.UserAuthRequest) (*model.UserAuthResponse, error) {
	hashedPassword, err := s.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	if err = s.repo.CreateUser(ctx, &user.User{
		ID:           uuid.New().String(),
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}); err != nil {
		return nil, err
	}

	return &model.UserAuthResponse{
		Message: userCreatedMsg,
	}, nil
}

func (s *impl) Login(ctx context.Context, req *model.UserAuthRequest) (*model.UserAuthResponse, error) {
	userData, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.PasswordHash), []byte(req.Password)); err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := generateTokens(userData.ID)
	if err != nil {
		return nil, err
	}

	err = s.repo.WithTransaction(ctx, func(txCtx context.Context) error {
		err = s.repo.CreateAccessToken(txCtx, &token.AccessToken{
			ID:        uuid.New().String(),
			UserID:    userData.ID,
			User:      *userData,
			TokenHash: accessToken,
			ExpiresAt: time.Now().Add(jwtExpiryTime),
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}

		err = s.repo.CreateRefreshToken(txCtx, &refreshtoken.RefreshToken{
			ID:               uuid.New().String(),
			UserID:           userData.ID,
			User:             *userData,
			RefreshTokenHash: refreshToken,
			ExpiresAt:        time.Now().Add(refreshExpiryTime),
			CreatedAt:        time.Now(),
		})
		if err != nil {
			return err
		}

		return nil
	})

	return &model.UserAuthResponse{
		Message: loginSuccessMsg,
		Token: &model.Tokens{
			AccessToken:        accessToken,
			AccessTokenExpiry:  time.Now().Add(jwtExpiryTime).Unix(),
			RefreshToken:       refreshToken,
			RefreshTokenExpiry: time.Now().Add(refreshExpiryTime).Unix(),
		},
	}, nil
}

var jwtKey = []byte("my_secret_key")

func generateTokens(userID string) (string, string, error) {
	accessTokenClaims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: time.Now().Add(jwtExpiryTime).Unix(),
	}

	refreshTokenClaims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: time.Now().Add(refreshExpiryTime).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *impl) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *impl) GetProfiles(ctx context.Context, userID string) (*model.UserProfileResponse, error) {
	userData, err := s.repo.GetUserByUserID(ctx, userID)
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

func (s *impl) RefreshToken(ctx context.Context, req *model.RefreshTokenRequest) (*model.UserAuthResponse, error) {
	refreshTokenData, err := s.repo.GetRefreshTokenByToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	if refreshTokenData.RevokedAt != nil || refreshTokenData.ExpiresAt.Before(time.Now()) {
		return nil, err
	}

	newAccessToken, newRefreshToken, err := generateTokens(refreshTokenData.UserID)
	if err != nil {
		return nil, err
	}

	err = s.repo.WithTransaction(ctx, func(txCtx context.Context) error {
		err = s.repo.CreateAccessToken(txCtx, &token.AccessToken{
			ID:        uuid.New().String(),
			UserID:    refreshTokenData.UserID,
			TokenHash: newAccessToken,
			ExpiresAt: time.Now().Add(jwtExpiryTime),
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}

		err = s.repo.CreateRefreshToken(txCtx, &refreshtoken.RefreshToken{
			ID:               uuid.New().String(),
			UserID:           refreshTokenData.UserID,
			RefreshTokenHash: newRefreshToken,
			ExpiresAt:        time.Now().Add(refreshExpiryTime),
			CreatedAt:        time.Now(),
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model.UserAuthResponse{
		Message: "Token refreshed successfully",
		Token: &model.Tokens{
			AccessToken:        newAccessToken,
			AccessTokenExpiry:  time.Now().Add(jwtExpiryTime).Unix(),
			RefreshToken:       newRefreshToken,
			RefreshTokenExpiry: time.Now().Add(refreshExpiryTime).Unix(),
		},
	}, nil
}

func (s *impl) RevokeToken(ctx context.Context, req *model.RevokeTokenRequest) error {
	refreshTokenData, err := s.repo.GetRefreshTokenByToken(ctx, req.AccessToken)
	if err != nil {
		return err
	}

	if refreshTokenData.RevokedAt != nil {
		return fmt.Errorf("refresh token is already revoked")
	}

	now := time.Now()
	refreshTokenData.RevokedAt = &now

	err = s.repo.WithTransaction(ctx, func(txCtx context.Context) error {
		err = s.repo.UpdateRefreshToken(txCtx, refreshTokenData)
		if err != nil {
			return err
		}

		err = s.repo.RevokeAccessTokenByUserID(txCtx, refreshTokenData.UserID)
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
