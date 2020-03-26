package district

import (
	"github.com/imjuanleonard/palu-covid/internal/district"
	"github.com/imjuanleonard/palu-covid/internal/handler/v1"
	"github.com/imjuanleonard/palu-covid/pkg/logger"
	"net/http"
)

type Handler struct {
	service district.Service
}

func NewDistrictHandler(service district.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) DistrictList(w http.ResponseWriter, r *http.Request) {
	d, err := h.service.List(r.Context())
	if err != nil {
		v1.NewErrorResponse(v1.CodeServerError, "Internal server error", "Internal server error").Write(w, http.StatusInternalServerError)
		logger.Errorf("[district.Handler.DistrictList] internal server error %v", err)
		return
	}
	v1.NewSuccessResponse(d).Write(w, http.StatusOK)
}
