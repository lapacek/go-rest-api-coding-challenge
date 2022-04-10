package data

import (
	"time"

	"github.com/lapacek/simple-api-example/internal/db"
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

func (r *Repository) BookTicket(booking Booking) error {

	return nil
}

func (r *Repository) GetDestinations() *[]Destination {

	return nil
}

func (r *Repository) GetBookings() *[]Booking {

	return nil
}

func (r *Repository) GetLaunches(from, to time.Time) *[]Launch {

	return nil
}

func (r *Repository) GetLaunch(date time.Time) *Launch {

	return nil
}