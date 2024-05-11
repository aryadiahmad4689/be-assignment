package module

import (
	"be-assesment/src/app"
	"be-assesment/src/modules/interfaces"
	"be-assesment/src/modules/payment_manager_services/endpoint"
	"be-assesment/src/modules/payment_manager_services/repository"
	transporthttp "be-assesment/src/modules/payment_manager_services/transport/http"
	"be-assesment/src/modules/payment_manager_services/usecase"

	"github.com/go-chi/chi"
)

type Module struct {
	usecase    usecase.UseCaseInterface
	endpoint   endpoint.EndpointInterface
	repository *repository.Repository
	httpRouter *chi.Mux
}

func Init(app *app.App) interfaces.ModuleInterface {
	var (
		repository = repository.Init(app)
		usecase    = usecase.Init(app, repository)
		endpoint   = endpoint.Init(app, usecase)
		httpRouter = transporthttp.Init(app, endpoint)
	)

	return &Module{
		repository: repository,
		usecase:    usecase,
		endpoint:   endpoint,
		httpRouter: httpRouter,
	}
}

func (module *Module) GetHttpRouter() *chi.Mux {
	return module.httpRouter
}
