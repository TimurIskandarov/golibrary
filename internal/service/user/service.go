package service

import (
	"context"

	"golibrary/internal/model"
	"golibrary/internal/repository/user"

	"go.uber.org/zap"
)

type Userer interface {
	List(ctx context.Context) ([]*model.User, error)
	Add(ctx context.Context, book model.User) error
}

type UserService struct {
	logger *zap.Logger
	repo   repo.Userer
}

func NewUserService(repo repo.Userer, logger *zap.Logger) *UserService {
	return &UserService{repo: repo, logger: logger}
}

func (us *UserService) List(ctx context.Context) ([]*model.User, error) {
	books, err := us.repo.List(ctx)
	if err != nil {
		us.logger.Error("error users list", zap.Error(err))
		return nil, err
	}
	return books, nil
}

func (us *UserService) Add(ctx context.Context, user model.User) error {
	if err := us.repo.Add(ctx, user); err != nil {
		us.logger.Error("error user add", zap.Error(err))
		return err
	}
	return nil
}
