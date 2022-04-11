package model

import "errors"

var (
	DestinationUnavailableError = errors.New("Destination is unavailable at the launch date")
	LaunchpadUsedBySpaceXError = errors.New("Launchpad is used by SpaceX company at the launch date")
	LaunchpadUnavailableError = errors.New("Launchpad is unavailable at the launch date")
	OutOfDateError = errors.New("Cannot book ticket to the history")

	RepositoryError = errors.New("Cannot retrieve data")
)
