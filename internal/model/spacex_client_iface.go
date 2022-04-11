package model

import (
	"context"
)

// allows polymorphic usage of the SpaceXClient object in the integrational tests
type SpaceXClientIface interface {
	GetLaunches(ctx context.Context) (*[]SpaceXLaunch, error)
}

