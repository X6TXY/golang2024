package service

import (
	"database/sql"
	"log"

	"github.com/x6txy/go2024/finalproject/model"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (us *UserService) CreateUser(user *model.User) (int, error) {
	if err := us.createUsersTableIfNotExists(); err != nil {
		log.Println("Error creating 'users' table:", err)
		return 0, err
	}

	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	var userID int
	err := us.DB.QueryRow(query, user.Name, user.Email).Scan(&userID)
	if err != nil {
		log.Println("Error creating user:", err)
		return 0, err
	}
	return userID, nil
}

func (us *UserService) GetUserByID(userID int) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	user := &model.User{}
	err := us.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Println("Error retrieving user:", err)
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetAllUsers() ([]*model.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := us.DB.Query(query)
	if err != nil {
		log.Println("Error retrieving users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Println("Error scanning user row:", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (us *UserService) createUsersTableIfNotExists() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255)
		);
	`
	_, err := us.DB.Exec(query)
	if err != nil {
		log.Println("Error creating 'users' table:", err)
		return err
	}
	return nil
}
