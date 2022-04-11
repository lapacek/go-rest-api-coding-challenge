package data

import (
	"context"
	"time"
)

// allows polymorphic usage of the Repository object in the integrational tests
type RepositoryIface interface {
	CreateLaunch(ctx context.Context, booking Booking) error
	BookTicket(ctx context.Context, booking Booking) error
	GetBookings(ctx context.Context) (*[]Booking, error)
	GetDestinations(ctx context.Context) (*[]Destination, error)
	GetLaunch(ctx context.Context, date time.Time) (*Launch, error)
	GetLaunches(ctx context.Context, from, to time.Time) (*[]Launch, error)
	GetLaunchpad(ctx context.Context, launchpadId string) (*LaunchPad, error)
}
