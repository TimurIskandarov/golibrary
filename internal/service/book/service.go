package service

import (
	"context"

	"golibrary/internal/model"
	repo "golibrary/internal/repository/book"

	"go.uber.org/zap"
)

type Booker interface {
	BookTake(ctx context.Context, userId, bookId int) error
	BookReturn(ctx context.Context, userId, bookId int) error
	BooksList(ctx context.Context) ([]*model.Book, error)
	BookAdd(ctx context.Context, book model.Book) error
}

type BookService struct {
	logger *zap.Logger
	repo repo.BookerRepository
}

func NewBookService(logger *zap.Logger) *BookService {
	return &BookService{logger: logger}
}

func (bs *BookService) BookTake(ctx context.Context, userId, bookId int) error {
	if err := bs.repo.BookTake(ctx, userId, bookId); err != nil {
		bs.logger.Error("error book take", zap.Error(err))
		return err
	}
	return nil
}

func (bs *BookService) BookReturn(ctx context.Context, userId, bookId int) error {
	if err := bs.repo.BookReturn(ctx, userId, bookId); err != nil {
		bs.logger.Error("error book return", zap.Error(err))
		return err
	}
	return nil
}

func (bs *BookService) BooksList(ctx context.Context) ([]*model.Book, error) {
	books, err := bs.repo.BooksList(ctx)
	if err != nil {
		bs.logger.Error("error books list", zap.Error(err))
		return nil, err
	}
	return books, nil
}

func (bs *BookService) BookAdd(ctx context.Context, book model.Book) error {
	if err := bs.repo.BookAdd(ctx, book); err != nil {
		bs.logger.Error("error book add", zap.Error(err))
		return err
	}
	return nil
}
