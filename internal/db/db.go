package db

import (
	"context"
	"fmt"
	"github.com/lapacek/simple-api-example/internal/common"

	"github.com/gookit/config/v2"
	"github.com/jackc/pgx/v4"
)

const (
	anonymizedPass = "XXXXX"
	pattern        = "postgres://%s:%s@%s:%s/%s"
)

type DB struct {
	connStr string
	opened  bool

	conf *config.Config
	conn *pgx.Conn
}

func NewDB(conf *config.Config) *DB {
	db := DB{}
	db.conf = conf

	return &db
}

func (db *DB) Open() bool {

	if !db.init() {
		fmt.Println("Cannot initialize database")

		return false
	}

	var err error
	db.conn, err = pgx.Connect(context.Background(), db.connStr)
	if err != nil {
		fmt.Printf("Cannot connect to database, err(%v)\n", err)

		return false
	}
	defer db.Close()

	db.opened = true
	fmt.Println("Database connected")

	return true
}

func (db *DB) Close() bool {

	if !db.opened {
		fmt.Println("Connection is already closed")
		return true
	}

	err := db.conn.Close(context.Background())
	if err != nil {
		fmt.Printf("Cannot close database, err(%v)\n", err)

		return false
	}
	db.opened = false

	return true
}

func (db *DB) init() bool {

	connStrWoPass := createConnStr(
		db.conf.String(common.PGHost),
		db.conf.String(common.PGPort),
		db.conf.String(common.PGUser),
		anonymizedPass,
		db.conf.String(common.PGName),
	)
	fmt.Printf("Connection string(%v)\n", connStrWoPass)

	db.connStr = createConnStr(
		db.conf.String(common.PGHost),
		db.conf.String(common.PGPort),
		db.conf.String(common.PGUser),
		db.conf.String(common.PGPass),
		db.conf.String(common.PGName),
	)

	return true
}

func (db *DB) ping() bool {

	fmt.Println("Pinging db...")

	if !db.opened {
		fmt.Println("Cannot ping closed connection")

		return false
	}

	err := db.conn.Ping(context.Background())
	if err != nil {
		fmt.Printf("Ping failed, err(%v)\n", err)

		return false
	}

	fmt.Println("Ping was successful")

	return true
}

// Builds correct connection string
//
// Pattern should be filled this way:
// "postgres://username:password@localhost:5432/database_name"
func createConnStr(host, port, user, pass, name string) string {

	return fmt.Sprintf(pattern, user, pass, host, port, name)
}
