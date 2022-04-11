package model_test

import (
	"context"
	"github.com/lapacek/simple-api-example/internal/model"
	data2 "github.com/lapacek/simple-api-example/internal/model/data"
	"testing"
	"time"
)

// TODO: launch dates should be generated

type RepositoryMock struct {
	Bookings *[]data2.Booking
	Destinations *[]data2.Destination
	Launch *data2.Launch
	Launches *[]data2.Launch
}

func (r RepositoryMock) CreateLaunch(ctx context.Context, booking data2.Booking) error {
	return nil
}

func (r RepositoryMock) BookTicket(ctx context.Context, booking data2.Booking) error {
	return nil
}

func (r RepositoryMock) GetBookings(ctx context.Context) (*[]data2.Booking, error) {
	return r.Bookings, nil
}

func (r RepositoryMock) GetDestinations(ctx context.Context) (*[]data2.Destination, error) {
	return r.Destinations, nil
}

func (r RepositoryMock) GetLaunch(ctx context.Context, date time.Time) (*data2.Launch, error) {
	return r.Launch, nil
}

func (r RepositoryMock) GetLaunches(ctx context.Context, from, to time.Time) (*[]data2.Launch, error) {
	return r.Launches, nil
}

var destinations = []data2.Destination{
	data2.Destination{ ID: "1", Name: "Mars"},
	data2.Destination{ ID: "2", Name: "Moon"},
	data2.Destination{ ID: "3", Name: "Pluto"},
	data2.Destination{ ID: "4", Name: "Asteroid Belt"},
	data2.Destination{ ID: "5", Name: "Europa"},
	data2.Destination{ ID: "6", Name: "Titan"},
	data2.Destination{ ID: "7", Name: "Ganymede"},
						}

// the target launch exists in the database, the model will book ticket for that launch
func Test_BookTicket(t *testing.T) {

	// all luanches in target week
	launches := make([]data2.Launch, 0)

	// target launch
	launch_1 := data2.Launch{}
	launch_1.ID = "1"
	launch_1.LaunchDate = "2022-04-13"
	launch_1.DestinationID = "3"
	launch_1.LaunchpadID = "2"

	launches = append(launches, launch_1)

	// previous launch
	launch_2 := data2.Launch{}
	launch_2.ID = "1"
	launch_2.LaunchDate = "2022-04-12"
	launch_2.DestinationID = "4"
	launch_2.LaunchpadID = "1"

	launches = append(launches, launch_2)

	// next launch
	launch_3 := data2.Launch{}
	launch_3.ID = "1"
	launch_3.LaunchDate = "2022-04-15"
	launch_3.DestinationID = "1"
	launch_3.LaunchpadID = "3"

	launches = append(launches, launch_3)

	repository := RepositoryMock{}
	repository.Launch = &launch_1
	repository.Launches = &launches
	repository.Destinations = &destinations

	// input booking data, points to the launch_1
	booking := model.Booking{}
	booking.LaunchDate = "2022-04-13"
	booking.DestinationID = "3"
	booking.LaunchpadID = "2"
	booking.FirstName = "Elon"
	booking.LastName = "Musk"
	booking.Birthday = "1971-06-28"
	booking.Birthday = "men"

	model := model.NewModel(repository)

	err := model.BookTicket(&booking)

	if err != nil {
		t.Errorf("Booking failed")
	}
}

// the target launch date of booking is lower than today, test doesn't need any launches
func Test_BookTicket_Fail_For_Incorrect_LuanchTime(t *testing.T) {

	// all luanches in target week
	launches := make([]data2.Launch, 0)

	repository := RepositoryMock{}
	repository.Launches = &launches
	repository.Destinations = &destinations

	// input booking data, points to the launch_1
	booking := model.Booking{}
	booking.LaunchDate = "2022-04-01"
	booking.DestinationID = "3"
	booking.LaunchpadID = "2"
	booking.FirstName = "Elon"
	booking.LastName = "Musk"
	booking.Birthday = "1971-06-28"
	booking.Birthday = "men"

	model := model.NewModel(repository)

	err := model.BookTicket(&booking)

	if err == nil {
		t.Errorf("Booking should failed because the launch date is out of date")
	}
}
