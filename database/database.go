package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang-solid-clean-architecture/config"
)

type Database interface {
	Execute(query string, args ...interface{}) error
	QueryRow(query string, args ...interface{}) (Row, error)
}

type Row interface {
	Scan(dest ...interface{}) error
}

type SQLRow struct {
	row *sql.Row
}

func (r SQLRow) Scan(dest ...interface{}) error {
	return r.row.Scan(dest...)
}

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(config config.DBConfig) (Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) Execute(query string, args ...interface{}) error {
	_, err := p.db.Exec(query, args...)
	return err
}

func (p *PostgresDB) QueryRow(query string, args ...interface{}) (Row, error) {
	return SQLRow{row: p.db.QueryRow(query, args...)}, nil
}
