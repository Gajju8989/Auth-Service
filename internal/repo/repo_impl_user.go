package repo

import (
	"context"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/user"
)

func (i *impl) CreateUser(ctx context.Context, userData *user.User) error {
	return i.db.
		WithContext(ctx).
		Create(userData).
		Error
}

func (i *impl) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	var userData user.User
	err := i.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&userData).
		Error

	return &userData, err
}

func (i *impl) GetUserByUserID(ctx context.Context, userID string) (*user.User, error) {
	var userData user.User
	err := i.db.
		WithContext(ctx).
		Where("id = ?", userID).
		First(&userData).
		Error

	return &userData, err
}
