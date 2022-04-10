package internal

import (
	"fmt"
	"github.com/lapacek/simple-api-example/internal/common"
	"time"

	"github.com/lapacek/simple-api-example/internal/data"
)

type Model struct {
	destinations *[]data.Destination

	repository *data.Repository
}

func NewModel(repository *data.Repository) *Model {
	model := Model{}
	model.repository = repository

	return &model
}

func (m *Model) Open() bool {

	return m.loadDestinations()
}

func (m *Model) Close() bool {

	return true
}

func (m *Model) loadDestinations() bool {

	dest := m.repository.GetDestinations()

	if dest == nil {
		fmt.Printf("Cannot load destinations")

		return false
	}

	m.destinations = dest

	return true
}

func (m *Model) SpaceXLaunches() *[]SpaceXLaunch {
	return nil
}

func (m *Model) AllBookings() *[]Booking {

	bookings := m.repository.GetBookings()
	results := make([]Booking, len(*bookings), 0)

	for _, b := range *bookings {
		results = append(results, Booking{Booking : b})
	}

	return &results
}

func (m *Model) BookTicket(booking *Booking) error {

	launchDate, err := time.Parse(common.DateLayout, booking.LaunchDate)
	if err != nil {
		return fmt.Errorf("Parsing failed, layout(%v), launchDate(%v)\n", common.DateLayout, booking.LaunchDate)
	}

	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	if launchDate.Before(today) {
		return fmt.Errorf("Launch date(%v) should be higher than/equal to today", booking.LaunchDate)
	}

	launch := m.repository.GetLaunch(launchDate)
	if launch != nil {
		if launch.LaunchpadID == booking.LaunchpadID {
			err := m.repository.BookTicket(booking.Booking)
			if err != nil {
				fmt.Printf("Booking failed, err(%v)\n", err)

				return fmt.Errorf("Booking failed, please try it later")
			}

			return nil
		}
	}

	launches := m.repository.GetLaunches(GetStartOfWeek(launchDate), GetEndOfWeek(launchDate))
	for _, l := range *launches {
		if l.DestinationID == booking.DestinationID {
			return fmt.Errorf(
				"We provide another launch to same destination at same week, please choose another date",
				)
		}
	}

	spacexLaunches := m.SpaceXLaunches()
	for _, sxl := range *spacexLaunches {
		if sxl.Upcomming {

			sxlLaunchDate, err := time.Parse(common.SpaceXDateTimeLayout, sxl.DateUTC)
			if err != nil {
				fmt.Printf(
					"Cannot parse SpaceX launch date(%v), layout(%v), launchID(%v), err(%v)\n",
					sxl.DateUTC,
					common.SpaceXDateTimeLayout,
					sxl.ID,
					err,
					)

				return fmt.Errorf("Unexpected error, please contact support")
			}

			if sxlLaunchDate == launchDate {
				if sxl.LaunchpadID == booking.LaunchpadID {
					return fmt.Errorf("There is no available launch on selected launchpad at the launch date")
				}
			}
		}
	}

	err = m.repository.BookTicket(booking.Booking)
	if err != nil {
		fmt.Printf("Cannot book ticket, err(%v)\n", err)

		return fmt.Errorf("Unexpected error, please contact support")
	}

	return nil
}

func (m *Model) DeleteBooking() {

}

func GetStartOfWeek(date time.Time) time.Time {

	weekday := date.Weekday()

	year, month, day := date.Date()
	dateWithDeletedTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return dateWithDeletedTime.Add(-1 * (time.Duration(weekday) - 1) * 24 * time.Hour)
}

func GetEndOfWeek(date time.Time) time.Time {

	weekday := date.Weekday()

	year, month, day := date.Date()
	dateWithDeletedTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	return dateWithDeletedTime.Add((7 - time.Duration(weekday)) * 24 * time.Hour)
}
