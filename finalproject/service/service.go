package service

import (
	"database/sql"
	"fmt"

	"github.com/x6txy/go2024/finalproject/model"
	"golang.org/x/crypto/bcrypt"
)

type BaseService struct {
	DB *sql.DB
}

func NewBaseService(db *sql.DB) *BaseService {
	return &BaseService{
		DB: db,
	}
}

func hashPassword(password []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	return hashedPassword, nil
}

type CommentService struct {
	*BaseService
}

type UserService struct {
	*BaseService
}

func NewCommentService(db *sql.DB) *CommentService {
	return &CommentService{
		BaseService: NewBaseService(db),
	}
}

func (cs *CommentService) CreateComment(comment *model.Comment) (int, error) {
	if err := cs.ensureCommentsTableExists(); err != nil {
		return 0, fmt.Errorf("failed to ensure comments table exists: %w", err)
	}

	const query = "INSERT INTO comments (text, date) VALUES ($1, $2) RETURNING id"
	var commentID int
	err := cs.DB.QueryRow(query, comment.Text, comment.Date).Scan(&commentID)
	if err != nil {
		return 0, fmt.Errorf("failed to create comment: %w", err)
	}

	return commentID, nil
}

func (cs *CommentService) GetCommentByID(commentID int) (*model.Comment, error) {
	const query = "SELECT id, text, date FROM comments WHERE id = $1"
	comment := &model.Comment{}
	err := cs.DB.QueryRow(query, commentID).Scan(&comment.ID, &comment.Text, &comment.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("comment not found with ID %d", commentID)
		}
		return nil, fmt.Errorf("failed to retrieve comment: %v", err)
	}

	return comment, nil
}

func (cs *CommentService) GetAllComments() ([]*model.Comment, error) {
	const query = "SELECT id, text, date FROM comments"
	rows, err := cs.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve comments: %v", err)
	}
	defer rows.Close()

	var comments []*model.Comment
	for rows.Next() {
		comment := &model.Comment{}
		err := rows.Scan(&comment.ID, &comment.Text, &comment.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment row: %v", err)
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (cs *CommentService) ensureCommentsTableExists() error {
	const query = `
		CREATE TABLE IF NOT EXISTS comments (
			id SERIAL PRIMARY KEY,
			text TEXT,
			date TIMESTAMP
		);
	`
	_, err := cs.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create 'comments' table: %w", err)
	}

	return nil
}

func (cs *CommentService) DeleteCommentByID(commentID int) error {
	const query = "DELETE FROM comments WHERE id = $1"
	_, err := cs.DB.Exec(query, commentID)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %v", err)
	}

	return nil
}

func (cs *CommentService) UpdateComment(comment *model.Comment) error {
	const query = "UPDATE comments SET text = $2, date = $3 WHERE id = $1"
	_, err := cs.DB.Exec(query, comment.ID, comment.Text, comment.Date)
	if err != nil {
		return fmt.Errorf("failed to update comment: %v", err)
	}

	return nil
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		BaseService: NewBaseService(db),
	}
}

func (us *UserService) ensureUsersTableExists() error {
	const query = `
		CREATE TABLE IF NOT EXISTS main_users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL
		);
	`
	_, err := us.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create 'main_users' table: %w", err)
	}

	return nil
}

func (us *UserService) GetUserByUsername(username string) (*model.User, error) {
	const query = "SELECT id, username, password_hash FROM main_users WHERE username = $1"
	user := &model.User{}
	err := us.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found with username %s", username)
		}
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	return user, nil
}

func (us *UserService) RegisterUser(user *model.User) (userID int, err error) {
	if err = us.ensureUsersTableExists(); err != nil {
		err = fmt.Errorf("failed to ensure main_users table exists: %w", err)
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		err = fmt.Errorf("failed to hash password: %w", err)
		return
	}

	const query = "INSERT INTO main_users (username, password_hash) VALUES ($1, $2) RETURNING id"
	err = us.DB.QueryRow(query, user.Username, hashedPassword).Scan(&userID)
	if err != nil {
		err = fmt.Errorf("failed to register user: %w", err)
		return
	}

	return userID, nil
}

func (us *UserService) GetAllUsers() ([]*model.User, error) {
	const query = "SELECT id, username, password_hash FROM main_users"
	rows, err := us.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %v", err)
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}
