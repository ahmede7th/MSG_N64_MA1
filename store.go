package main

import (
    "database/sql"
)


type Store interface {
    CreateStuff(stuff *Stuff) error
    GetStuff() ([]*Stuff, error)
}

type dbStore struct {
    db *sql.DB
}

func (store *dbStore) CreateStuff(stuff *Stuff) error {
	_, err := store.db.Query("INSERT INTO stuffs(name, price) VALUES ($1,$2)", stuff.Name, stuff.Price)
	return err
}

func (store *dbStore) GetStuffs() ([]*Stuff, error) {
	rows, err := store.db.Query("SELECT name, price from stuffs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stuffs := []*Stuff{}
	for rows.Next() {
		stuff := &Stuff{}
		if err := rows.Scan(&stuff.Name, &stuff.Price); err != nil {
			return nil, err
		}
		stuffs = append(stuffs, stuff)
	}
	return stuffs, nil
}

var store Store


func InitStore(s Store) {
	store = s
}
