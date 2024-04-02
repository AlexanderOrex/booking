package booking

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"applicationDesignTest/internal/pkg/model"
	"applicationDesignTest/internal/pkg/provider/orders"
)

type BookingProvider interface {
	CreateOrders(ctx context.Context, orders []model.Order) error
}

type Handler struct {
	bookingProvider BookingProvider
}

func NewHandler(bookingProvider BookingProvider) *Handler {
	return &Handler{
		bookingProvider: bookingProvider,
	}
}

type Order struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (h *Handler) CreateOrders(w http.ResponseWriter, r *http.Request) {
	var newOrder Order

	json.NewDecoder(r.Body).Decode(&newOrder)

	byDate, err := decodeToOrdersByDate(newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.bookingProvider.CreateOrders(r.Context(), byDate); err != nil {
		status := http.StatusInternalServerError
		if err == orders.ErrValidation {
			status = http.StatusBadRequest
		} else if err == orders.ErrBusy {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}

func decodeToOrdersByDate(newOrder Order) ([]model.Order, error) {
	if newOrder.From.After(newOrder.To) || newOrder.From.Equal(newOrder.To) {
		return nil, errors.New("некорректные даты заказа")
	}

	result := []model.Order{}
	current := newOrder.From
	// Проходим по всем датам и добавляем их в слайс
	for current.Before(newOrder.To) {
		result = append(result, model.Order{
			HotelID:   model.HotelID(newOrder.HotelID),
			RoomID:    model.RoomID(newOrder.RoomID),
			UserEmail: newOrder.UserEmail,
			Date:      model.DateFromTime(current),
		})
		fmt.Println(current)
		current = current.AddDate(0, 0, 1) // Добавляем один день к текущей дате
	}

	return result, nil
}
