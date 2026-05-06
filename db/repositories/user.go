package db

import (
	model "authService/models"
	"database/sql"
	"fmt"
)
type UserRepository interface{
	Create() (error)
	GetById(id string) (error, *model.User)
}

type UserRepositoryImpl struct{
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository{
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) Create() (error){
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?);"
	_, err := u.db.Exec(query,"test_user", "test@gail.com", "test@123")
	if err != nil {
        fmt.Println("Error executing query:", err)
        return err
    }
    fmt.Println("user created successfully")
	return nil
}

func (u *UserRepositoryImpl) GetById(id string) (error, *model.User){
	//setting up the query
	query := "SELECT * FROM users WHERE id=UUID_TO_BIN(?)"
	//execute the query
	row := u.db.QueryRow(query, id)
	//processing the query and creating object out of row data
	user := &model.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	if err == sql.ErrNoRows {
		fmt.Println("No data fount for the id", id)
		return err, user
	}
	fmt.Println("user fetched successfully")
	return nil, user
}