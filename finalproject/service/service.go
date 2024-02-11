package service

import (
	"database/sql"
	"fmt"

	"github.com/x6txy/go2024/finalproject/model"
)

type CommentService struct {
	DB *sql.DB
}

func NewCommentService(db *sql.DB) *CommentService {
	return &CommentService{
		DB: db,
	}
}

func (cs *CommentService) CreateComment(comment *model.Comment) (int, error) {
	if err := cs.ensureCommentsTableExists(); err != nil {
		return 0, fmt.Errorf("failed to ensure comments table exists: %v", err)
	}

	const query = "INSERT INTO comments (text, date) VALUES ($1, $2) RETURNING id"
	var commentID int
	err := cs.DB.QueryRow(query, comment.Text, comment.Date).Scan(&commentID)
	if err != nil {
		return 0, fmt.Errorf("failed to create comment: %v", err)
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
		return fmt.Errorf("failed to create 'comments' table: %v", err)
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
