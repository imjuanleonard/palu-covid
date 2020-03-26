package district

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("district not found")
)

type Service interface {
	List(ctx context.Context) ([]District, error)
	GetByID(ctx context.Context, id string) (*District, error)
	Update(ctx context.Context, district *District) error
}

type service struct {
	store *store
}

func NewService(db *sqlx.DB) Service {
	return &service{
		store: &store{db},
	}
}

func (svc *service) List(ctx context.Context) ([]District, error) {
	districs, err := svc.store.list(ctx)
	if err != nil {
		return []District{}, errors.Wrap(err, "[district.List]")
	}
	return districs, nil
}

func (svc *service) GetByID(ctx context.Context, id string) (*District, error) {
	districs, err := svc.store.findByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return districs, nil
}

func (svc *service) Update(ctx context.Context, district *District) error {
	return svc.store.update(ctx, district)
}
