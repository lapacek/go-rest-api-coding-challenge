package data

import (
	"github.com/lapacek/simple-api-example/internal"
	"time"
)

type Repository struct {
	db *internal.DB
}

func NewRepository(db *internal.DB) *Repository {
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

func (r *Repository) GetDestinations() {

}

func (r *Repository) GetLaunches(from, to *time.Time) {

}

func (r *Repository) GetTickets() {

}