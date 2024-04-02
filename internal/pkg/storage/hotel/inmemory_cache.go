package hotel

import (
	"applicationDesignTest/internal/pkg/model"
	"context"
	"fmt"
	"sync"
	"time"
)

type RoomAvailability struct {
	HotelID string
	RoomID  string
	Date    time.Time
	Quota   int
}

type InmemoryCache struct {
	availability map[string]uint32

	// TODO: mutex by hotelID
	mx sync.Mutex
}

func NewCache() *InmemoryCache {
	return &InmemoryCache{
		availability: map[string]uint32{
			key(model.HotelID("reddison"), model.RoomID("lux"), model.Date(date(2025, 1, 1))): 1,
			key(model.HotelID("reddison"), model.RoomID("lux"), model.Date(date(2025, 1, 2))): 1,
			key(model.HotelID("reddison"), model.RoomID("lux"), model.Date(date(2025, 1, 3))): 2,
			key(model.HotelID("reddison"), model.RoomID("lux"), model.Date(date(2025, 1, 4))): 1,
			key(model.HotelID("reddison"), model.RoomID("lux"), model.Date(date(2025, 1, 5))): 0,
		},
	}
}

func (c *InmemoryCache) TryToReserveHotels(ctx context.Context, orders []model.Order) (bool, error) {
	c.mx.Lock()
	defer c.mx.Unlock()

	grouppedOrders := make(map[string]uint32, 1)
	for _, order := range orders {
		grouppedOrders[key(order.HotelID, order.RoomID, order.Date)]++
	}

	// check
	for key, quota := range grouppedOrders {
		if c.availability[key] < quota {
			return false, nil
		}
	}

	// apply
	for key, quota := range grouppedOrders {
		c.availability[key] = c.availability[key] - quota
	}

	return true, nil
}

// Вспомогательная функция. В проде будет удалена
func date(year, month, day int) model.Date {
	return model.Date(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}

func key(hotelID model.HotelID, roomID model.RoomID, date model.Date) string {
	return fmt.Sprintf("%v:%v%v", hotelID, roomID, date)
}
