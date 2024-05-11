package handler

	import (
		"be-assesment/src/app"
		"be-assesment/src/modules/payment_manager_services/endpoint"
	)
	
	type Handler struct {
		app      *app.App
		endpoint endpoint.EndpointInterface
	}
	
	func InitHandler(app *app.App, endpoint endpoint.EndpointInterface) *Handler {
		var handler = &Handler{
			app:      app,
			endpoint: endpoint,
		}
		return handler
	}	
	