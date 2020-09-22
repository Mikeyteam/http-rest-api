package sqlstore

import (
	"database/sql"
	"github.com/Mikeyteam/http-rest-api/internal/store"
	_ "github.com/lib/pq" //Пока импортнули анонимно, т.е без применения в коде
)

type Store struct {
	db *sql.DB
	userRepository *UserRepository
}

//Конструктор
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//В дальнейшем можно будет вызвать так store.User().Create()
func (s *Store) User() store.UserRepository {
    if s.userRepository != nil {
    	return s.userRepository
	}
	//Если не сущ. инициализируем репозиторий
	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
