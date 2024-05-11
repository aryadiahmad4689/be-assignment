package handler

import (
	"encoding/json"
	"net/http"

	"be-assesment/src/modules/payment_manager_services/endpoint"
	"be-assesment/src/modules/payment_manager_services/transport/http/response"
)

func (handler *Handler) Send(w http.ResponseWriter, r *http.Request) {
	var (
		data   endpoint.SendReq
		err    error
		ctx    = r.Context()
		result interface{}
	)

	// Decode the request body into data
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "Failed Send", err.Error())
		return
	}

	// Call the Send
	result, err = handler.endpoint.Send(ctx, data)
	if err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "Failed Send", err.Error())
		return
	}

	response.SendSuccessResponse(w, http.StatusOK, "Success", result)
}
