package service

import (
	"database/sql"
	"fmt"
	"time"


	"github.com/x6txy/go2024/finalproject/model"

)

type Database struct {
	DB *sql.DB
}

func NewDatabaseService(db *sql.DB) *Database {
	return &Database{
		DB: db,
	}
}

// PostService

type PostService struct {
	*Database
}

func NewPostService(db *sql.DB) *PostService {
	return &PostService{
		Database: NewDatabaseService(db),
	}
}

func (cs *PostService) CreatePost(post *model.Post) (int, error) {
	if err := cs.ensurePostTableExists(); err != nil {
		return 0, fmt.Errorf("failed to ensure posts table exists: %w", err)
	}
	const query = "INSERT INTO posts (text, date, update_date) VALUES ($1, $2, $3) RETURNING id"
	var postID int
	err := cs.DB.QueryRow(query, post.Text, post.Date, post.Update_date).Scan(&postID)
	if err != nil {
		return 0, fmt.Errorf("failed to create post: %w", err)
	}

	return postID, nil
}

func (cs *PostService) GetAllPosts() ([]*model.Post, error) {
	const query = "SELECT * FROM posts"
	rows, err := cs.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve posts: %v", err)
	}
	defer rows.Close()

	var posts []*model.Post
	for rows.Next() {
		post := &model.Post{}
		err := rows.Scan(&post.ID, &post.Text, &post.Date, &post.Update_date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan post row: %v", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (cs *PostService) GetPostByID(postID int) (*model.Post, error) {
	const query = "SELECT * FROM posts WHERE id = $1"
	post := &model.Post{}
	err := cs.DB.QueryRow(query, postID).Scan(&post.ID, &post.Text, &post.Date, &post.Update_date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post not found with ID %d", postID)
		}
		return nil, fmt.Errorf("failed to retrieve post: %v", err)
	}

	return post, nil
}

func (cs *PostService) DeletePostById(postID int) error {
	const query = "DELETE FROM posts WHERE id = $1"
	_, err := cs.DB.Exec(query, postID)
	if err != nil {
		return fmt.Errorf("failed to delete post: %v", err)
	}
	return nil
}

func (cs *PostService) UpdatePost(postID int, newText string) error {
	const query = "UPDATE posts SET text = $1, update_date = $2 WHERE id = $3"
	_, err := cs.DB.Exec(query, newText, time.Now().Truncate(time.Minute), postID)
	if err != nil {
		return fmt.Errorf("failed to update post: %v", err)
	}
	return nil
}

func (cs *PostService) ensurePostTableExists() error {
	const query = `
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			text TEXT,
			date TIMESTAMP,
			update_date TIMESTAMP
		);
	`
	_, err := cs.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create 'posts' table: %w", err)
	}

	return nil
}


//CommentService

type CommentService struct {
	*Database
}

func NewCommentService(db *sql.DB) *CommentService {
	return &CommentService{
		Database: NewDatabaseService(db),
	}
}




func (cm *CommentService) ensureCommentTableExists() error{
	const query = `
	CREATE TABLE IF NOT EXISTS comments (
		id SERIAL PRIMARY KEY,
		post_id INT,
		date TIMESTAMP,
		content TEXT, 
		likes INT,
		user_id INT,
		FOREIGN KEY (post_id) REFERENCES posts(id)
	);`

	_, err := cm.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create 'comments' table: %w", err)
	}

	return nil
}