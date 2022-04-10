package internal

import (
	"time"
)

type Model struct {
	db *DB
}

func NewModel(db *DB) *Model {
	model := Model{}
	model.db = db

	return &model
}

func (m *Model) Open() bool {

	m.loadDestinations()

	return true
}

func (m *Model) Close() bool {

	return true
}

func (m *Model) loadDestinations() {

}

func (m *Model) loadSpaceXLaunches() {

}

func (m *Model) AllBookings() {

}

func (m *Model) BookTicket(booking *Booking) {

}

func (m *Model) DeleteBooking() {

}

func GetStartOfWeek(date time.Time) string {

	weekday := date.Weekday()

	year, month, day := date.Date()
	dateWithDeletedTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	result := dateWithDeletedTime.Add(-1 * (time.Duration(weekday) - 1) * 24 * time.Hour)

	return result.Format(DateLayout)
}

func GetEndOfWeek(date time.Time) string {

	weekday := date.Weekday()

	year, month, day := date.Date()
	dateWithDeletedTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	result := dateWithDeletedTime.Add((7 - time.Duration(weekday)) * 24 * time.Hour)

	return result.Format(DateLayout)
}
