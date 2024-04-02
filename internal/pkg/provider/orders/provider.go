package orders

import (
	"applicationDesignTest/internal/pkg/model"
	"context"
	"errors"
	"fmt"
)

type Validation interface {
	ValidateOrders(ctx context.Context, orders []model.Order) (bool, error)
}

type Hotels interface {
	TryToReserveHotels(ctx context.Context, orders []model.Order) (bool, error)
}

type BookingProcessor interface {
	ProcessOrders(ctx context.Context, orders []model.Order) error
}

type OrdersProvider struct {
	validation       Validation
	hotels           Hotels
	bookingProcessor BookingProcessor
}

func NewOrdersProvider(
	validation Validation,
	hotels Hotels,
	bookingProcessor BookingProcessor,
) *OrdersProvider {
	return &OrdersProvider{
		validation:       validation,
		hotels:           hotels,
		bookingProcessor: bookingProcessor,
	}
}

var (
	ErrValidation = errors.New("ошибка валидации даты")
	ErrBusy       = errors.New("нет свободных номеров на данную дату")
)

func (p *OrdersProvider) CreateOrders(ctx context.Context, orders []model.Order) error {
	ok, err := p.validation.ValidateOrders(ctx, orders)
	if err != nil {
		return err
	}
	if !ok {
		return ErrValidation
	}

	ok, err = p.hotels.TryToReserveHotels(ctx, orders)
	if err != nil {
		return err
	}
	if !ok {
		return ErrBusy
	}

	// TODO: outbox
	if err = p.bookingProcessor.ProcessOrders(ctx, orders); err != nil {
		fmt.Println("Не удалось обработать заказ")
	}

	return nil
}
