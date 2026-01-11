package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	repository "github.com/shindeshubhamm/go-ecomm/internal/adapters/postgresql/sqlc"
)

type ProductService interface {
	ListProducts(ctx context.Context) ([]repository.Product, error)
	GetProductById(ctx context.Context, id pgtype.UUID) (repository.Product, error)
}

type productService struct {
	repo repository.Querier
}

func NewProductService(repo repository.Querier) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) ListProducts(ctx context.Context) ([]repository.Product, error) {
	products, err := s.repo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) GetProductById(ctx context.Context, id pgtype.UUID) (repository.Product, error) {
	product, err := s.repo.FindProductById(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return repository.Product{}, errors.New("not found")
		}
		return repository.Product{}, err
	}
	return product, nil
}
