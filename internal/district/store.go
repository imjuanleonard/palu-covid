package district

import (
	"context"
	"github.com/imjuanleonard/palu-covid/config"
	"github.com/imjuanleonard/palu-covid/pkg/db"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type store struct {
	db *sqlx.DB
}

func (s *store) districtList(ctx context.Context) ([]District, error) {
	getQuery := "select id, nama, odp, pdp, positif, negatif, meninggal, selesai_pengawasan, dalam_pengawasan, selesai_pemantauan, dalam_pemantauan, created_at, updated_at from kabupaten limit 100"
	district := []District{}
	queryFunction := func(ctx context.Context) error {
		return db.Get().SelectContext(ctx, &district, getQuery)
	}
	if err := db.WithTimeout(ctx, config.Database.ReadTimeoutSecond, queryFunction); err != nil {
		return district, errors.Wrap(err, "[district.listKabupaten]")
	}
	return district, nil
}
