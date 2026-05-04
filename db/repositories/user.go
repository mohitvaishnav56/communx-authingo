package db

import "database/sql"
type UserRepository interface{
	// Create() error
}

type UserRepositoryImpl struct{
	db *sql.DB
}

func (u *UserRepositoryImpl) Create() error{
	return nil
}

func NewUserRepository() UserRepository{
	return &UserRepositoryImpl{
		db: nil,
	}
}