package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //...
)

type Store struct{
	config *Config
	db *sql.DB
	userrepository *UserRepository
}

func New(config *Config) *Store{
	return &Store{
		config:config,
	}
}

func(s *Store) Open() error{

	//fmt.Print("Открывает бд")

	db, err := sql.Open("postgres",s.config.DatabaseUrl)
	fmt.Print(db)
	if err != nil{
		
		return err
	}
	if err := db.Ping(); err != nil{
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository{
	if s.userrepository != nil{
		return s.userrepository
	}
	s.userrepository = &UserRepository{
		store: s,
	}
	return s.userrepository
}