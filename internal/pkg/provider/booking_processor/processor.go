package booking_processor

import (
	"applicationDesignTest/internal/pkg/model"
	"context"
	"fmt"
)

type BookingProcessor struct {
}

func NewBookingProcessor() *BookingProcessor {
	return &BookingProcessor{}
}

func (v *BookingProcessor) ProcessOrders(ctx context.Context, orders []model.Order) error {
	for _, order := range orders {
		fmt.Println(fmt.Sprintf("order created by user %s", order.UserEmail), order)
	}

	return nil
}
