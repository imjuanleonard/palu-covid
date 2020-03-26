package district

import (
	"github.com/gorilla/mux"
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

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	districtID := vars["district_id"]
	d, err := h.service.GetByID(ctx, districtID)
	if err != nil {
		switch err {
		case district.ErrDistrictNotFound:
			{
				v1.NewErrorResponse(v1.CodeInvalidRequest, "District ID not found", "Bad Request").Write(w, http.StatusBadRequest)
				logger.Warnf("[district.Handler.GetByID] district with id = %s does not exist", districtID)
				return
			}
		default:
			v1.NewErrorResponse(v1.CodeServerError, "Internal Server Error", "Internal Server Error").Write(w, http.StatusInternalServerError)
			logger.Errorf("[ChannelHandler.GetByID] internal server error %v", err)
			return
		}
	}
	v1.NewSuccessResponse(d).Write(w, http.StatusOK)
}
