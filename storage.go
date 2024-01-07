package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgreStore struct {
	db *sql.DB
}

func NewPostgreStore() (*PostgreStore, error) {
	connStr := "user=postgres dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreStore {
		db: db,
	}, nil
}