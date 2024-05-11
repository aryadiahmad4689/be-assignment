package http

import (
	"be-assesment/src/app"
	"be-assesment/src/modules/account_manager_services/endpoint"
	"be-assesment/src/modules/account_manager_services/transport/http/handler"

	"github.com/go-chi/chi"
)

func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {
	var (
		router = chi.NewRouter()
		h      = handler.InitHandler(app, endpoint)
	)

	router.Post("/sign-up", h.SignUp)
	router.Post("/sign-in", h.SignIn)
	router.Get("/get-all-user-account", h.GetAllUsersWithPayments)

	return router
}
