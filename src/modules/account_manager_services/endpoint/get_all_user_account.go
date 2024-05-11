package endpoint

import (
	"context"
)

func (ed *Endpoint) GetAllUsersWithPayments(ctx context.Context) (interface{}, error) {
	data, err := ed.usecase.GetAllUsersWithPayments(ctx)
	if err != nil {
		return RepNull{}, err
	}
	return data, nil
}
