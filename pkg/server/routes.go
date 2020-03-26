package server

import (
	"github.com/gorilla/mux"
	h "github.com/imjuanleonard/palu-covid/internal/handler"
	"github.com/imjuanleonard/palu-covid/pkg/middleware"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r = r.SkipClean(true)
	r.Use(middleware.Recover())
	r.HandleFunc("/ping", h.PingHandler).Methods(http.MethodGet)
	return r
}

func NewRouter(handler *h.Handler) *mux.Router {
	r := newRouter()

	r.HandleFunc("/kabupaten", handler.District.DistrictList).Methods(http.MethodGet)
	return r
}
