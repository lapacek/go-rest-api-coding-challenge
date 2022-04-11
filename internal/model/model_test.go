package model_test

import (
	"context"
	"testing"
	"time"

	"github.com/lapacek/simple-api-example/internal/model"
	"github.com/lapacek/simple-api-example/internal/model/data"
)

// TODO: launch dates should be generated

var destinations = []data.Destination{
	data.Destination{ ID: "1", Name: "Mars"},
	data.Destination{ ID: "2", Name: "Moon"},
	data.Destination{ ID: "3", Name: "Pluto"},
	data.Destination{ ID: "4", Name: "Asteroid Belt"},
	data.Destination{ ID: "5", Name: "Europa"},
	data.Destination{ ID: "6", Name: "Titan"},
	data.Destination{ ID: "7", Name: "Ganymede"},
}

var spacexLaunches = []model.SpaceXLaunch{
	model.SpaceXLaunch{ ID: "61eefaa89eb1064137a1bd73", LaunchpadID: "5e9e4502f509094188566f88", DateUTC: "2022-04-16T15:04:05.000Z", Upcomming: true},
}

type SpaceXClientMock struct {
	Launches *[]model.SpaceXLaunch
}

func (c SpaceXClientMock) GetLaunches(ctx context.Context) (*[]model.SpaceXLaunch, error) {
	return c.Launches, nil
}

type RepositoryMock struct {
	Bookings *[]data.Booking
	Destinations *[]data.Destination
	Launch *data.Launch
	Launches *[]data.Launch
	LaunchPad *data.LaunchPad
}

func (r RepositoryMock) CreateLaunch(ctx context.Context, booking data.Booking) error {
	return nil
}

func (r RepositoryMock) BookTicket(ctx context.Context, booking data.Booking) error {
	return nil
}

func (r RepositoryMock) GetBookings(ctx context.Context) (*[]data.Booking, error) {
	return r.Bookings, nil
}

func (r RepositoryMock) GetDestinations(ctx context.Context) (*[]data.Destination, error) {
	return r.Destinations, nil
}

func (r RepositoryMock) GetLaunch(ctx context.Context, date time.Time) (*data.Launch, error) {
	return r.Launch, nil
}

func (r RepositoryMock) GetLaunches(ctx context.Context, from, to time.Time) (*[]data.Launch, error) {
	return r.Launches, nil
}

func (r RepositoryMock) GetLaunchpad(ctx context.Context, launchpadId string) (*data.LaunchPad, error) {
	return r.LaunchPad, nil
}

// the target launch exists in the database, the model will book ticket for that launch
func Test_BookTicket(t *testing.T) {

	// all luanches in target week
	launches := make([]data.Launch, 0)

	// target launch
	launch_1 := data.Launch{}
	launch_1.ID = "1"
	launch_1.LaunchDate = "2022-04-13"
	launch_1.DestinationID = "3"
	launch_1.LaunchpadID = "2"

	launches = append(launches, launch_1)

	// previous launch
	launch_2 := data.Launch{}
	launch_2.ID = "1"
	launch_2.LaunchDate = "2022-04-12"
	launch_2.DestinationID = "4"
	launch_2.LaunchpadID = "1"

	launches = append(launches, launch_2)

	// next launch
	launch_3 := data.Launch{}
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

	api := model.NewModel(repository, SpaceXClientMock{})

	err := api.BookTicket(&booking)
	if err != nil {
		t.Errorf("Booking failed, err(%v)", err)
	}
}

// the target launch date of booking is lower than today, test doesn't need any launches
func Test_BookTicket_Fail_For_Incorrect_LuanchTime(t *testing.T) {

	// all luanches in target week
	launches := make([]data.Launch, 0)

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

	api := model.NewModel(repository, SpaceXClientMock{})

	err := api.BookTicket(&booking)
	if err != nil {
		if err == model.OutOfDateError {
			return
		}

		t.Errorf("Booking failed, err(%v)", err)
	}

	t.Errorf("Booked, but the booking should failed because the launch date is out of date")
}

// we don't provide launch to the same destination more than once per week
func Test_BookTicket_Fail_For_Unavailable_Destination(t *testing.T) {

	// all luanches in target week
	launches := make([]data.Launch, 0)

	launch_1 := data.Launch{}
	launch_1.ID = "1"
	launch_1.LaunchDate = "2022-04-13"
	launch_1.DestinationID = "3"
	launch_1.LaunchpadID = "2"

	launches = append(launches, launch_1)

	launch_2 := data.Launch{}
	launch_2.ID = "1"
	launch_2.LaunchDate = "2022-04-12"
	launch_2.DestinationID = "4"
	launch_2.LaunchpadID = "1"

	launches = append(launches, launch_2)

	launch_3 := data.Launch{}
	launch_3.ID = "1"
	launch_3.LaunchDate = "2022-04-15"
	launch_3.DestinationID = "1"
	launch_3.LaunchpadID = "3"

	launches = append(launches, launch_3)

	repository := RepositoryMock{}
	repository.Launches = &launches
	repository.Destinations = &destinations

	booking := model.Booking{}
	booking.LaunchDate = "2022-04-16"
	booking.DestinationID = "3"
	booking.LaunchpadID = "2"
	booking.FirstName = "Elon"
	booking.LastName = "Musk"
	booking.Birthday = "1971-06-28"
	booking.Birthday = "men"

	api := model.NewModel(repository, SpaceXClientMock{})

	err := api.BookTicket(&booking)
	if err != nil {
		if err == model.DestinationUnavailableError {
			return
		}

		t.Errorf("Booking failed, err(%v)", err)
	}

	t.Errorf("Booked, but the booking should failed because the destination is used in existing launch")
}

// we don't provide launch from the same launchpad as spacex at the same day
func Test_BookTicket_Fail_For_Launchpad_Reserved_By_SpaceX(t *testing.T) {

	// all luanches in target week
	launches := make([]data.Launch, 0)

	launch_1 := data.Launch{}
	launch_1.ID = "1"
	launch_1.LaunchDate = "2022-04-13"
	launch_1.DestinationID = "3"
	launch_1.LaunchpadID = "2"

	launches = append(launches, launch_1)

	launch_2 := data.Launch{}
	launch_2.ID = "1"
	launch_2.LaunchDate = "2022-04-12"
	launch_2.DestinationID = "4"
	launch_2.LaunchpadID = "1"

	launches = append(launches, launch_2)

	launch_3 := data.Launch{}
	launch_3.ID = "1"
	launch_3.LaunchDate = "2022-04-15"
	launch_3.DestinationID = "1"
	launch_3.LaunchpadID = "3"

	launches = append(launches, launch_3)

	repository := RepositoryMock{}
	repository.Launches = &launches
	repository.LaunchPad = &data.LaunchPad{ID: "2", IDSpaceX: "5e9e4502f509094188566f88"}
	repository.Destinations = &destinations

	booking := model.Booking{}
	booking.LaunchDate = "2022-04-16"
	booking.DestinationID = "6"
	booking.LaunchpadID = "2"
	booking.FirstName = "Elon"
	booking.LastName = "Musk"
	booking.Birthday = "1971-06-28"
	booking.Birthday = "men"

	spacexClient := SpaceXClientMock{}
	spacexClient.Launches = &spacexLaunches

	api := model.NewModel(repository, spacexClient)

	err := api.BookTicket(&booking)
	if err != nil {
		if err == model.LaunchpadUsedBySpaceXError {
			return
		}

		t.Errorf("Booking failed, err(%v)", err)
	}

	t.Errorf("Booked, but the booking should failed because the launchpad is reserved by SpaceX")
}

func Test_BookTicket_And_Create_Launch(t *testing.T) {

	// all luanches in target week
	launches := make([]data.Launch, 0)

	launch_1 := data.Launch{}
	launch_1.ID = "1"
	launch_1.LaunchDate = "2022-04-13"
	launch_1.DestinationID = "3"
	launch_1.LaunchpadID = "2"

	launches = append(launches, launch_1)

	launch_2 := data.Launch{}
	launch_2.ID = "1"
	launch_2.LaunchDate = "2022-04-12"
	launch_2.DestinationID = "4"
	launch_2.LaunchpadID = "1"

	launches = append(launches, launch_2)

	launch_3 := data.Launch{}
	launch_3.ID = "1"
	launch_3.LaunchDate = "2022-04-15"
	launch_3.DestinationID = "1"
	launch_3.LaunchpadID = "3"

	launches = append(launches, launch_3)

	repository := RepositoryMock{}
	repository.Launches = &launches
	repository.Destinations = &destinations

	booking := model.Booking{}
	booking.LaunchDate = "2022-04-17"
	booking.DestinationID = "5"
	booking.LaunchpadID = "2"
	booking.FirstName = "Elon"
	booking.LastName = "Musk"
	booking.Birthday = "1971-06-28"
	booking.Birthday = "men"

	spacexClient := SpaceXClientMock{}
	spacexClient.Launches = &spacexLaunches

	api := model.NewModel(repository, spacexClient)

	err := api.BookTicket(&booking)
	if err != nil {
		t.Errorf("Booking failed, err(%v)", err)
	}
}
