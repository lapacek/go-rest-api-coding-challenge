package data

import (
	"context"
	"fmt"
	"github.com/lapacek/simple-api-example/internal/common"
	"github.com/lapacek/simple-api-example/internal/lib/db"
	"time"
)

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	repo := Repository{}
	repo.db = db

	return &repo
}

func (r *Repository) Open() bool {

	return true
}

func (r *Repository) Close() bool {

	return true
}

func (r Repository) CreateLaunch(ctx context.Context, booking Booking) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		fmt.Println("Cannot begin transaction")

		return err
	}

	tx.Exec(ctx, INSERT_LAUNCH, booking.LaunchID, booking.DestinationID, booking.LaunchDate)
	tx.Commit(ctx)

	return nil
}

func (r Repository) BookTicket(ctx context.Context, booking Booking) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		fmt.Println("Cannot begin transaction")

		return err
	}

	tx.Exec(ctx, INSERT_TICKET, booking.LaunchID, booking.FirstName, booking.LastName, booking.Gender, booking.Birthday)
	tx.Commit(ctx)

	return nil
}

func (r Repository) GetDestinations(ctx context.Context) (*[]Destination, error) {

	rows, err := r.db.Query(ctx, SELECT_DESTINATIONS)
	if err != nil {
		fmt.Println("Cannot retrieve destinations")

		return nil, err
	}

	results := make([]Destination, 0)

	for rows.Next() {
		var item Destination

		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			fmt.Println("Cannot scan destination from db row")

			return nil, err
		}

		results = append(results, item)
	}

	return &results, nil
}

func (r Repository) GetBookings(ctx context.Context) (*[]Booking, error) {

	rows, err := r.db.Query(ctx, SELECT_BOOKINGS)
	if err != nil {
		fmt.Println("Cannot retrieve bookings")

		return nil, err
	}

	results := make([]Booking, 0)

	for rows.Next() {

		var item Booking

		err := rows.Scan(&item.FirstName, &item.LastName, &item.Gender,
						&item.Birthday, &item.LaunchpadID, &item.DestinationID,
						&item.LaunchDate,
						)
		if err != nil {
			fmt.Println("Cannot scan booking from db row")

			return nil, err
		}

		results = append(results, item)
	}

	return &results, nil
}

func (r Repository) GetLaunches(ctx context.Context, from, to time.Time) (*[]Launch, error) {

	start := from.Format(common.DateLayout)
	end := to.Format(common.DateLayout)

	rows, err := r.db.Query(ctx, SELECT_LAUNCHES, start, end)
	if err != nil {
		fmt.Println("Cannot retrieve launches")

		return nil, err
	}

	results := make([]Launch, 0)

	for rows.Next() {
		var item Launch

		err := rows.Scan(&item.ID, &item.LaunchpadID, &item.DestinationID, &item.LaunchDate)
		if err != nil {
			fmt.Println("Cannot scan launch from db row")

			return nil, err
		}

		results = append(results, item)
	}

	return &results, nil
}

func (r Repository) GetLaunch(ctx context.Context, date time.Time) (*Launch, error) {

	launchDate := date.Format(common.DateLayout)

	row := r.db.QueryRow(ctx, SELECT_LAUNCH, launchDate)

	var item Launch

	err := row.Scan(&item.ID, &item.LaunchpadID, &item.DestinationID, &item.LaunchDate)
	if err != nil {
		fmt.Println("Cannot scan launch from db row")

		return nil, err
	}

	return &item, nil
}

func (r Repository) GetLaunchpad(ctx context.Context, launchpadId int) (*LaunchPad, error) {

	row := r.db.QueryRow(ctx, SELECT_LAUNCHPAD, launchpadId)

	var item LaunchPad

	err := row.Scan(&item.ID, &item.IDSpaceX)
	if err != nil {
		fmt.Println("Cannot scan launchpad from db row")

		return nil, err
	}

	return &item, nil
}