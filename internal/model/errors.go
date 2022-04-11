package model

import "errors"

var (
	BookingError                = errors.New("Booking failed, please contact support")
	DestinationUnavailableError = errors.New("Destination is unavailable at the launch date, " +
													"We provide another launch to same destination at same week, " +
													"please choose another date")
	LaunchpadUsedBySpaceXError = errors.New("Launchpad is used by SpaceX company at the launch date")
	LaunchpadUnavailableError  = errors.New("Launchpad is unavailable at the launch date")
	OutOfDateError             = errors.New("Cannot book ticket to the history")

	RepositoryError = errors.New("Cannot retrieve data")
)
