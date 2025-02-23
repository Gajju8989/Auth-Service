package service

import (
	"context"
	"github.com/google/uuid"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/user"
	"github/com/Gajju8989/Auth_Service/internal/service/model"
	"net/http"
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
		if isMySQLDuplicateEntryError(err) {
			return nil, &GenericResponse{
				StatusCode: http.StatusConflict,
				Message:    emailConflictErrMessage,
			}
		}

		return nil, err
	}

	return &model.UserAuthResponse{
		Message: userCreatedMsg,
	}, nil
}
