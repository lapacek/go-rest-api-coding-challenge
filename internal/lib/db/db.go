package db

import (
	"context"
	"fmt"

	"github.com/lapacek/simple-api-example/internal/common"

	"github.com/gookit/config/v2"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	anonymizedPass = "XXXXX"
	pattern        = "postgres://%s:%s@%s:%s/%s"
)

type Rows interface {
	pgx.Rows
}

type Row interface {
	pgx.Row
}

type Tx interface {
	pgx.Tx
}

type DB struct {
	connStr string
	opened  bool

	conf *config.Config
	connPool *pgxpool.Pool
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
	db.connPool, err = pgxpool.Connect(context.Background(), db.connStr)
	if err != nil {
		fmt.Printf("Cannot connect to database, err(%v)\n", err)

		return false
	}

	db.opened = true
	fmt.Println("Database connected")

	return true
}

func (db *DB) Close() bool {

	if !db.opened {
		fmt.Println("Connection is already closed")
		return true
	}

	db.connPool.Close()
	db.opened = false

	return true
}

func (db *DB) Query(ctx context.Context, sql string, args ...interface{}) (Rows, error) {
	return db.connPool.Query(ctx, sql, args...)
}

func (db *DB) QueryRow(ctx context.Context, sql string, args ...interface{}) Row {
	return db.connPool.QueryRow(ctx, sql, args...)
}

func (db *DB) Begin(ctx context.Context) (Tx, error) {
	return db.connPool.Begin(ctx)
}

func (db *DB) Rollback(ctx context.Context, tx Tx) error {
	return tx.Rollback(ctx)
}

func (db *DB) Commit(ctx context.Context, tx Tx) error {
	return tx.Commit(ctx)
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

	err := db.connPool.Ping(context.Background())
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
