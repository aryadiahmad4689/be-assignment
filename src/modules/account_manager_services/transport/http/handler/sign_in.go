package handler

import (
	"encoding/json"
	"net/http"

	"be-assesment/src/modules/account_manager_services/endpoint"
	"be-assesment/src/modules/account_manager_services/transport/http/response"
)

func (handler *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var (
		data   endpoint.SignInRequest
		err    error
		ctx    = r.Context()
		result interface{}
	)

	// Decode the request body into data
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "Failed SignIn", err.Error())
		return
	}

	// Call the SignIn endpoint
	result, err = handler.endpoint.SignIn(ctx, data)
	if err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "Failed SignIn", err.Error())
		return
	}

	response.SendSuccessResponse(w, http.StatusOK, "Success", result)
}
