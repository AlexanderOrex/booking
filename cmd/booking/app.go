package booking

import (
	"context"

	booking_handler "applicationDesignTest/internal/app/booking"
	booking_provider "applicationDesignTest/internal/pkg/provider/booking_processor"
	orders_provideer "applicationDesignTest/internal/pkg/provider/orders"
	validator_provider "applicationDesignTest/internal/pkg/provider/validator"
	hotel_cache "applicationDesignTest/internal/pkg/storage/hotel"
)

type Storage struct {
	hotels *hotel_cache.InmemoryCache
}

type Providers struct {
	booking          *booking_provider.BookingProcessor
	bookingProcessor *booking_provider.BookingProcessor
	orders           *orders_provideer.OrdersProvider
	validator        *validator_provider.ValidationProvider
}

type Handlers struct {
	Booking *booking_handler.Handler
}

type App struct {
	providers Providers
	Handlers  Handlers
}

func InitApp(ctx context.Context) *App {
	storages := Storage{
		hotels: hotel_cache.NewCache(),
	}

	bookingProcessor := booking_provider.NewBookingProcessor()
	validation := validator_provider.NewValidationProvider()

	providers := Providers{
		bookingProcessor: bookingProcessor,
		orders:           orders_provideer.NewOrdersProvider(validation, storages.hotels, bookingProcessor),
		validator:        validation,
	}
	handlers := Handlers{
		Booking: booking_handler.NewHandler(providers.orders),
	}

	return &App{
		Handlers: handlers,
	}

}
