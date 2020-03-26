package district

import (
	"errors"
	"time"
)

var ErrDistrictNotFound = errors.New("district does not found")

type District struct {
	ID                   int       `db:"id" json:"id"`
	Name                 string    `db:"nama" json:"nama"`
	ODP                  int       `db:"odp" json:"odp" `
	PDP                  int       `db:"pdp" json:"pdp"`
	Positive             int       `db:"positif" json:"positif"`
	Negative             int       `db:"negatif" json:"negatif"`
	PassAway             int       `db:"meninggal" json:"meninggal"`
	CompletedSupervision int       `db:"selesai_pengawasan" json:"selesai_pengawasan"`
	UnderSupervision     int       `db:"dalam_pengawasan" json:"dalam_pengawasan"`
	CompletedObservation int       `db:"selesai_pemantauan" json:"selesai_pemantauan"`
	UnderObservation     int       `db:"dalam_pemantauan" json:"dalam_pemantauan"`
	CreatedAt            time.Time `db:"created_at" json:"-"`
	UpdatedAt            time.Time `db:"updated_at" json:"-"`
}
