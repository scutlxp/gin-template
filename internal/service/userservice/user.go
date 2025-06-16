package userservice

import (
	"context"
	"gin-project/internal/types/usertypes"
)

type UserService struct {
}

func Get() UserService {
	return UserService{}
}

func (u UserService) GetUsers(ctx context.Context) ([]usertypes.User, error) {
	return nil, nil
}
