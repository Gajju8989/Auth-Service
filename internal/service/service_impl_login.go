package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github/com/Gajju8989/Auth_Service/internal/service/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func (s *impl) Login(ctx context.Context, req *model.UserAuthRequest) (*model.UserAuthResponse, error) {
	userData, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &GenericResponse{
				StatusCode: http.StatusNotFound,
				Message:    userNotFoundMsg,
			}
		}

		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.PasswordHash), []byte(req.Password)); err != nil {
		return nil, err
	}

	accessTokenUUID := uuid.New().String()
	accessToken, err := s.generateAccessToken(accessTokenUUID)
	if err != nil {
		return nil, err
	}

	refreshTokenUUID := uuid.New().String()
	refreshToken, err := s.generateRefreshToken(refreshTokenUUID)
	if err != nil {
		return nil, err
	}

	err = s.createTokens(ctx, accessTokenUUID, refreshTokenUUID, userData.ID)
	if err != nil {
		return nil, err
	}

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
