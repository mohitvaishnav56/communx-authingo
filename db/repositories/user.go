package db

import (
	model "authService/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(username string, email string, hashedPassword string) error
	GetById(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]*model.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) error {
	newUUID := uuid.New().String()
	query := "INSERT INTO users (id, username, email, password) VALUES (UUID_TO_BIN(?), ?, ?, ?);"
	_, err := u.db.Exec(query, newUUID, username, email, hashedPassword)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return err
	}
	fmt.Println("user created successfully")
	return nil
}

func (u *UserRepositoryImpl) GetById(id string) (*model.User, error) {
	//setting up the query
	query := "SELECT BIN_TO_UUID(id) AS id, username, email, IFNULL(created_at, ''), IFNULL(updated_at, '') FROM users WHERE id=UUID_TO_BIN(?)"
	//execute the query
	row := u.db.QueryRow(query, id)
	//processing the query and creating object out of row data
	user := &model.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Created_at, &user.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No data found for the id", id)
		} else {
			fmt.Println("Error scanning user:", err)
		}
		return nil, err
	}

	fmt.Println("user fetched successfully")
	return user, nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	query := "SELECT BIN_TO_UUID(id), username, email, password, IFNULL(created_at, ''), IFNULL(updated_at, '') FROM users WHERE email=?"
	row := u.db.QueryRow(query, email)
	user := &model.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) GetAll() ([]*model.User, error) {
	//query
	query := "SELECT BIN_TO_UUID(id) AS id, username, email FROM users"
	//executing the query
	rows, err := u.db.Query(query)
	if err != nil {
		fmt.Println("Query error: %v", err)
		return nil, err
	}
	defer rows.Close()
	var users []*model.User
	for rows.Next() {
		u := &model.User{}
		if err := rows.Scan(&u.Id, &u.Username, &u.Email); err != nil {
			fmt.Println("Scan error: %v", err)
		}
		users = append(users, u)
	}
	// fmt.Println(users)
	return users, nil
}
