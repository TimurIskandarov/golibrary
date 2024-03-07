package service

import (
	"context"

	"golibrary/internal/model"
	"golibrary/internal/book/repository"

	"go.uber.org/zap"
)

type Booker interface {
	Take(ctx context.Context, userId, bookId int) (*model.Book, error)
	Return(ctx context.Context, userId, bookId int) (*model.Book, error)
	List(ctx context.Context) ([]*model.Book, error)
	Add(ctx context.Context, book model.Book) (int, error)
}

type BookService struct {
	logger *zap.Logger
	repo repo.Booker
}

func NewBookService(repo repo.Booker, logger *zap.Logger) *BookService {
	return &BookService{repo: repo, logger: logger}
}

func (bs *BookService) Take(ctx context.Context, userId, bookId int) (*model.Book, error) {
	book, err := bs.repo.Take(ctx, userId, bookId)
	if err != nil {
		bs.logger.Error("error book take", zap.Error(err))
		return nil, err
	}
	return book, nil
}

func (bs *BookService) Return(ctx context.Context, userId, bookId int) (*model.Book, error) {
	book, err := bs.repo.Return(ctx, userId, bookId)
	if err != nil {
		bs.logger.Error("error book return", zap.Error(err))
		return nil, err
	}
	return book, nil
}

func (bs *BookService) List(ctx context.Context) ([]*model.Book, error) {
	books, err := bs.repo.List(ctx)
	if err != nil {
		bs.logger.Error("error books list", zap.Error(err))
		return nil, err
	}
	return books, nil
}

func (bs *BookService) Add(ctx context.Context, book model.Book) (int, error) {
	bookId, err := bs.repo.Add(ctx, book)
	if err != nil {
		bs.logger.Error("error book add", zap.Error(err))
		return 0, err
	}
	return bookId, nil
}
