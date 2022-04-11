package data

import (
	"context"
	"time"
)

type RepositoryIface interface {
	CreateLaunch(ctx context.Context, booking Booking) error
	BookTicket(ctx context.Context, booking Booking) error
	GetBookings(ctx context.Context) (*[]Booking, error)
	GetDestinations(ctx context.Context) (*[]Destination, error)
	GetLaunch(ctx context.Context, date time.Time) (*Launch, error)
	GetLaunches(ctx context.Context, from, to time.Time) (*[]Launch, error)
}
