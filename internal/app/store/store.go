package store

import (
	"database/sql"
	_ "github.com/lib/pq"//..
)

//Store...
type Store struct{
	config *Config
	db *sql.DB
	cityRepository *CityRepo
	ticketRepostiry *TicketRepo
	passengerRepository *PassengerRepo
}

func New(config *Config) *Store{
	return &Store{
		config:config,
	}
}
//Open...
func (s *Store) Open() error{
	connStr:="host=localhost port=5432 user=procol dbname=airport sslmode=disable"
	db,err:=sql.Open("postgres",connStr)
	if err!=nil{
		return err
	}
	if err:=db.Ping();err!=nil{
		return err
	}
	s.db=db
	return nil
}

func (s *Store) Close(){
	s.db.Close()
}



func (s *Store) Ticket() *TicketRepo{
	if s.cityRepository!=nil {
		return s.ticketRepostiry
	}
	s.ticketRepostiry = &TicketRepo{
		store:s,

	}
	return s.ticketRepostiry
}

func (s *Store) Passenger() *PassengerRepo{
	if s.passengerRepository!=nil{
		return s.passengerRepository
	}
	s.passengerRepository=&PassengerRepo{
		store: s,
	}
	return s.passengerRepository
}
//City...
func (s *Store) City() *CityRepo{
	if s.cityRepository!=nil{
		return s.cityRepository
	}
	s.cityRepository=&CityRepo{
		store:s,
	}
	return s.cityRepository
}
//for example store.City().Create()