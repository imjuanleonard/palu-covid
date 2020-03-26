package district

import (
	"context"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type Service interface {
	List(ctx context.Context) ([]District, error)
}

type service struct {
	db *store
}

func NewService(db *sqlx.DB) Service {
	return &service{
		db: &store{db},
	}
}

func (s *service) List(ctx context.Context) ([]District, error) {
	districs, err := s.db.districtList(ctx)
	if err != nil {
		return []District{}, errors.Wrap(err, "[district.List]")
	}
	return districs, nil
}
