package httpsrv

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type API struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *API {
	return &API{
		logger: logger,
	}
}

func (a *API) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Post("/orders", a.createOrder)
		r.Get("/orders", a.getOrders)
		r.Get("/orders/id/status", a.getStatus)
	})
}
