package internal

import (
	"fmt"
	"github.com/lapacek/simple-api-example/internal/lib/db"
	model2 "github.com/lapacek/simple-api-example/internal/model"

	"github.com/gookit/config/v2"
)

type Server struct {
	conf  *config.Config
	db    *db.DB
	model *model2.Model
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

	s.run()
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

	if ! s.perpareServer() {
		fmt.Println("Cannot prepare server")

		return false
	}

	return true
}

func (s *Server) Close() bool {

	return true
}

func (s *Server) prepareServer() bool {

	return true
}

func (s *Server) run() {

	for true {
	}

}
