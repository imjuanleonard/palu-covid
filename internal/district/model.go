package district

import (
	"errors"
	"time"
)

var ErrDistrictNotFound = errors.New("district does not found")

type District struct {
	ID                   int       `db:"id" json:"id" validate:"required"`
	Name                 string    `db:"nama" json:"nama" validate:"required"`
	ODP                  int       `db:"odp" json:"odp"  validate:"required"`
	PDP                  int       `db:"pdp" json:"pdp" validate:"required"`
	Positive             int       `db:"positif" json:"positif" validate:"required"`
	Negative             int       `db:"negatif" json:"negatif" validate:"required"`
	PassAway             int       `db:"meninggal" json:"meninggal" validate:"required"`
	CompletedSupervision int       `db:"selesai_pengawasan" json:"selesai_pengawasan" validate:"required"`
	UnderSupervision     int       `db:"dalam_pengawasan" json:"dalam_pengawasan" validate:"required"`
	CompletedObservation int       `db:"selesai_pemantauan" json:"selesai_pemantauan" validate:"required"`
	UnderObservation     int       `db:"dalam_pemantauan" json:"dalam_pemantauan" validate:"required"`
	CreatedAt            time.Time `db:"created_at" json:"-"`
	UpdatedAt            time.Time `db:"updated_at" json:"-"`
}
