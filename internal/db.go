package internal

import (
	"database/sql"
	"fmt"

	"github.com/gookit/config/v2"
	_ "github.com/lib/pq"
)

const (
	anonymizedPass = "XXXXX"
	driver         = "postgres"
	pattern        = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

type DB struct {
	connStr string
	opened  bool

	conf *config.Config
	conn *sql.DB
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
	db.conn, err = sql.Open(driver, db.connStr)
	if err != nil {
		fmt.Printf("Cannot open database, err(%v)\n", err)

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

	err := db.conn.Close()
	if err != nil {
		fmt.Printf("Cannot close database, err(%v)\n", err)

		return false
	}
	db.opened = false

	return true
}

func (db *DB) init() bool {

	connStrWoPass := createConnStr(
		db.conf.String(PGHost),
		db.conf.String(PGPort),
		db.conf.String(PGUser),
		anonymizedPass,
		db.conf.String(PGName),
	)
	fmt.Printf("Connection string(%v)\n", connStrWoPass)

	db.connStr = createConnStr(
		db.conf.String(PGHost),
		db.conf.String(PGPort),
		db.conf.String(PGUser),
		db.conf.String(PGPass),
		db.conf.String(PGName),
	)

	return true
}

func (db *DB) ping() bool {

	if !db.opened {
		fmt.Println("Cannot ping closed connection")

		return false
	}

	err := db.conn.Ping()
	if err != nil {
		fmt.Printf("Ping failed, err(%v)\n", err)

		return false
	}

	return true
}

func createConnStr(host, port, user, pass, name string) string {

	return fmt.Sprintf(pattern, host, port, user, pass, name)
}
