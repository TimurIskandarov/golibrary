package service

import (
	"context"

	"golibrary/internal/model"
	repo "golibrary/internal/repository/author"

	"go.uber.org/zap"
)

type Authorer interface {
	AuthorsTop(ctx context.Context) ([]*model.Author, error)
	AuthorsList(ctx context.Context) ([]*model.Author, error)
	AddAuthor(ctx context.Context, author model.Author) error
}

type AuthorService struct {
	logger *zap.Logger
	repo   repo.AuthorerRepository
}

func NewAuthorService(logger *zap.Logger) *AuthorService {
	return &AuthorService{logger: logger}
}

func (as *AuthorService) AuthorsTop(ctx context.Context) ([]*model.Author, error) {
	authors, err := as.repo.AuthorsTop(ctx)
	if err != nil {
		as.logger.Error("error authors top", zap.Error(err))
		return nil, err
	}
	return authors, nil
}

func (as *AuthorService) AuthorsList(ctx context.Context) ([]*model.Author, error) {
	authors, err := as.repo.AuthorsList(ctx)
	if err != nil {
		as.logger.Error("error authors list", zap.Error(err))
		return nil, err
	}
	return authors, nil
}

func (as *AuthorService) AddAuthor(ctx context.Context, author model.Author) error {
	if err := as.repo.AddAuthor(ctx, author.Name, author.BirthDate); err != nil {
		as.logger.Error("error add author", zap.Error(err))
		return err
	}
	return nil
}
