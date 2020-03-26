package handler

import (
	"github.com/imjuanleonard/palu-covid/internal/district"
	districtHandler "github.com/imjuanleonard/palu-covid/internal/handler/v1/district"
	"github.com/imjuanleonard/palu-covid/pkg/db"
)

type Handler struct {
	District *districtHandler.Handler
}

func NewHandler() *Handler {
	database := db.Get()
	districtService := district.NewService(database)
	return &Handler{
		District: districtHandler.NewDistrictHandler(districtService),
	}
}
