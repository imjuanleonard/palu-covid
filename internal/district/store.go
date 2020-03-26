package district

import (
	"context"
	"database/sql"
	"github.com/imjuanleonard/palu-covid/config"
	"github.com/imjuanleonard/palu-covid/pkg/db"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strconv"
)

var ErrInvalidID = errors.New("invalid district ID")

type store struct {
	db *sqlx.DB
}

func (s *store) list(ctx context.Context) ([]District, error) {
	const listQuery = "select id, nama, odp, pdp, positif, negatif, meninggal, selesai_pengawasan, dalam_pengawasan, selesai_pemantauan, dalam_pemantauan, created_at, updated_at from kabupaten limit 100"

	district := []District{}
	queryFunction := func(ctx context.Context) error {
		return db.Get().SelectContext(ctx, &district, listQuery)
	}

	if err := db.WithTimeout(ctx, config.Database.ReadTimeoutSecond, queryFunction); err != nil {
		return district, errors.Wrap(err, "[district.listKabupaten]")
	}

	return district, nil
}

func (s *store) findByID(ctx context.Context, id string) (*District, error) {
	const findByIDQuery = "select id, nama, odp, pdp, positif, negatif, meninggal, selesai_pengawasan, dalam_pengawasan, selesai_pemantauan, dalam_pemantauan, created_at, updated_at from kabupaten where id=$1 limit 1"
	var idNumber int
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.Wrap(ErrInvalidID, "[district.store.findByID]")
	}

	var district District
	queryFunction := func(ctx context.Context) error {
		return s.db.GetContext(ctx, &district, findByIDQuery, idNumber)
	}

	if err := db.WithTimeout(ctx, config.Database.ReadTimeoutSecond, queryFunction); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrDistrictNotFound
		}
		return nil, errors.Wrap(err, "[district.store.findByID]")
	}

	return &district, nil
}
