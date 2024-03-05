package service

import (
	"context"

	"golibrary/internal/model"
	"golibrary/internal/repository/book"

	"go.uber.org/zap"
)

type Booker interface {
	BookTake(ctx context.Context, userId, bookId int) (*model.Book, error)
	BookReturn(ctx context.Context, userId, bookId int) (*model.Book, error)
	BooksList(ctx context.Context) ([]*model.Book, error)
	BookAdd(ctx context.Context, book model.Book) error
}

type BookService struct {
	logger *zap.Logger
	repo repo.BookerRepository
}

func NewBookService(repo repo.BookerRepository, logger *zap.Logger) *BookService {
	return &BookService{repo: repo, logger: logger}
}

func (bs *BookService) BookTake(ctx context.Context, userId, bookId int) (*model.Book, error) {
	book, err := bs.repo.BookTake(ctx, userId, bookId)
	if err != nil {
		bs.logger.Error("error book take", zap.Error(err))
		return nil, err
	}
	return book, nil
}

func (bs *BookService) BookReturn(ctx context.Context, userId, bookId int) (*model.Book, error) {
	book, err := bs.repo.BookReturn(ctx, userId, bookId)
	if err != nil {
		bs.logger.Error("error book return", zap.Error(err))
		return nil, err
	}
	return book, nil
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
