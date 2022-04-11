package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lapacek/simple-api-example/internal/handler"
	"github.com/lapacek/simple-api-example/internal/lib/db"
	"github.com/lapacek/simple-api-example/internal/model"
	"github.com/lapacek/simple-api-example/internal/model/data"

	"github.com/gookit/config/v2"
	"github.com/gorilla/mux"
)

type Server struct {
	conf  *config.Config
	db    *db.DB
	httpServer *http.Server
	model *model.Model
	router *mux.Router
}

func NewServer(conf *config.Config) *Server {
	s := Server{}
	s.conf = conf

	return &s
}

func (s *Server) Run() {

	if !s.Open() {
		panic("Cannot start server")
	}
	defer func() {
		if !s.Close() {
			panic("Cannot stop server")
		}
	}()

	s.serve()
}

func (s *Server) Open() bool {

	failed := false

	fmt.Println("Server is starting...")
	defer func() {
		if !failed {
			fmt.Println("Server started")
		}
	}()

	s.db = db.NewDB(s.conf)

	if !s.db.Open() {
		fmt.Println("Cannot open database")
		failed = true

		return false
	}

	repository := data.NewRepository(s.db)
	if !repository.Open() {
		fmt.Println("Cannot open data layer")
		failed = true

		return false
	}

	spacexClient := model.NewSpaceXClient()
	if !spacexClient.Open() {
		fmt.Println("Cannot open spacex client")
		failed = true

		return false
	}

	s.model = model.NewModel(*repository, *spacexClient)
	if !s.model.Open() {
		fmt.Println("Cannot open model")
		failed = true

		return false
	}

	if ! s.prepareServer() {
		fmt.Println("Cannot prepare server")

		return false
	}

	return true
}

func (s *Server) Close() bool {

	return true
}

func (s *Server) prepareServer() bool {

	s.router = mux.NewRouter().StrictSlash(true)

	s.setHandlers()

	s.httpServer = &http.Server{
		Addr: ":8080",
		Handler: s.router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return true
}

func (s *Server) setHandlers() {
	s.router.Path("/booking").Methods("GET").HandlerFunc(s.getBookings)
	s.router.Path("/bookings").Methods("POST").HandlerFunc(s.createBooking)
}

func (s *Server) getBookings(w http.ResponseWriter, r *http.Request) {
	handler.AllBookings(s.model, w)
}

func (s *Server) createBooking(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) serve() {

	err := s.httpServer.ListenAndServe()
	if err != nil {
		fmt.Printf("Http server failed, err(%v)", err)
	}
}
