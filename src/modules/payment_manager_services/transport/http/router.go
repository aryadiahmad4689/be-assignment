package http

import (
	"be-assesment/src/app"
	"be-assesment/src/middleware"
	"be-assesment/src/modules/payment_manager_services/endpoint"
	"be-assesment/src/modules/payment_manager_services/transport/http/handler"

	"github.com/go-chi/chi"
)

func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {
	var (
		router = chi.NewRouter()
		h      = handler.InitHandler(app, endpoint)
	)

	router.Post("/send", middleware.Auth(h.Send))
	router.Post("/widthraw", middleware.Auth(h.Withdraw))

	return router
}
