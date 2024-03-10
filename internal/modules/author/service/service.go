package service

import (
	"context"
	
	"golibrary/internal/model"
	"golibrary/internal/modules/author/repository"

	"go.uber.org/zap"
)

type Authorer interface {
	Top(ctx context.Context) ([]*model.Author, error)
	List(ctx context.Context) ([]*model.Author, error)
	Add(ctx context.Context, author model.Author) error
}

type AuthorService struct {
	logger *zap.Logger
	repo   repo.Authorer
}

func NewAuthorService(repo repo.Authorer, logger *zap.Logger) *AuthorService {
	return &AuthorService{repo: repo, logger: logger}
}

func (as *AuthorService) Top(ctx context.Context) ([]*model.Author, error) {
	authors, err := as.repo.Top(ctx)
	if err != nil {
		as.logger.Error("error authors top", zap.Error(err))
		return nil, err
	}
	return authors, nil
}

func (as *AuthorService) List(ctx context.Context) ([]*model.Author, error) {
	authors, err := as.repo.List(ctx)
	if err != nil {
		as.logger.Error("error authors list", zap.Error(err))
		return nil, err
	}
	return authors, nil
}

func (as *AuthorService) Add(ctx context.Context, author model.Author) error {
	if err := as.repo.Add(ctx, author); err != nil {
		as.logger.Error("error add author", zap.Error(err))
		return err
	}
	return nil
}
