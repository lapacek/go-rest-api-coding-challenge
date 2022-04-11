package model

import "errors"

var (
	BookingError                = errors.New("Booking failed, please contact support")
	DestinationUnavailableError = errors.New("Launch from the requested launchpad to the target destination is unavailable, " +
													"You can try to book another launch to same destination from another launchpad this week, " +
													"or you can try to choose another week")
	LaunchpadUsedBySpaceXError = errors.New("Launchpad is used by SpaceX company at the launch date")
	OutOfDateError             = errors.New("Cannot book ticket to the history")

	DataLayerError = errors.New("Cannot retrieve data, please contact support")

	LogicError = errors.New("Something gone wrong, please contact support")
)
