package validator

import (
	"applicationDesignTest/internal/pkg/model"
	"context"
	"time"
)

type ValidationProvider struct {
}

func NewValidationProvider() *ValidationProvider {
	return &ValidationProvider{}
}

func (v *ValidationProvider) ValidateOrders(ctx context.Context, orders []model.Order) (bool, error) {
	today := model.DateFromTime(time.Now())
	for _, order := range orders {
		if order.Date.Before(today) {
			return false, nil
		}
	}

	return true, nil
}
