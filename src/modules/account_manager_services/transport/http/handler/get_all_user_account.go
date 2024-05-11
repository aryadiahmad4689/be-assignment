package handler

import (
	"be-assesment/src/modules/account_manager_services/transport/http/response"
	"net/http"
)

func (handler *Handler) GetAllUsersWithPayments(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		ctx    = r.Context()
		result interface{}
	)
	// Call the SignIn endpoint
	result, err = handler.endpoint.GetAllUsersWithPayments(ctx)
	if err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "Failed GetAllUsersWithPayments", err.Error())
		return
	}

	response.SendSuccessResponse(w, http.StatusOK, "Success", result)
}
