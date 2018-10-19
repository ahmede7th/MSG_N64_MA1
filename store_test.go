package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite

	store *dbStore
	db    *sql.DB
}

func (s *StoreSuite) SetupSuite() {

	connString := "dbname=<your test db name> sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {

	_, err := s.db.Query("DELETE FROM stuffs")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateStuff() {
	s.store.CreateStuff(&Stuff{
		Name:    "test name",
		Species: "test species",
	})

	res, err := s.db.Query(`SELECT COUNT(*) FROM stuffs WHERE name='test name' AND SPECIES='test species'`)
	if err != nil {
		s.T().Fatal(err)
	}

	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}

func (s *StoreSuite) TestGetStuff() {
	_, err := s.db.Query(`INSERT INTO stuffs (species, name) VALUES('stuff','name')`)
	if err != nil {
		s.T().Fatal(err)
	}

	stuffs, err := s.store.GetStuffs()
	if err != nil {
		s.T().Fatal(err)
	}

	nStuffs := len(stuffs)
	if nStuffs != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", nStuffs)
	}

	expectedStuff := Stuff{"stuff", "name"}
	if *stuffs[0] != expectedStuff {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedStuff, *stuffs[0])
	}
}
