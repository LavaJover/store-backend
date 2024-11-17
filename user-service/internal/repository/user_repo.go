package user_repo

import (
	"database/sql"
	"log"

	user "github.com/LavaJover/store-backend/user-service/proto"
)

type UserRepository struct{
	DB *sql.DB
}

func NewUserRepository (db *sql.DB) *UserRepository{
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(newUser *user.CreateUserRequest) error{
	query := "INSERT INTO users (name, email, hashed_password) VALUES ($1, $2, $3)"
	query_row := r.DB.QueryRow(query, newUser.Name, newUser.Email, newUser.RawPassword)
	log.Println(query_row)
	return nil
}

func (r *UserRepository) GetUser(userRequest *user.UserRequest) error{
	query := "SELECT * FROM users WHERE id==$1"
	_ = r.DB.QueryRow(query, userRequest.Id)
	return nil
}