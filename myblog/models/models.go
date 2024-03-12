package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	// Followers int       `json:"followers"`
	Posts     []Post    `json:"posts" gorm:"foreignKey:UserID"`
	Comments  []Comment `json:"comments"`
}

type Post struct {
	gorm.Model
	Content    string    `json:"content"`
	UserID     uint      `json:"user_id"`
	User       User      `json:"-" gorm:"foreignKey:UserID"`
	Comments   []Comment `json:"comments" gorm:"foreignKey:PostID"`
	LikesCount int       `json:"likes_count" gorm:"-"`
}

type Comment struct {
	gorm.Model
	Content    string `json:"content"`
	UserID     uint   `json:"user_id"`
	PostID     uint   `json:"post_id"`
	LikesCount int    `json:"likes_count" gorm:"-"`
}

type PostLike struct {
	UserID uint `gorm:"index:idx_user_post,uniqueIndex" json:"user_id"`
	PostID uint `gorm:"index:idx_user_post,uniqueIndex" json:"post_id"`
}

type CommentLike struct {
	UserID    uint `gorm:"index:idx_user_comment,uniqueIndex" json:"user_id"`
	CommentID uint `gorm:"index:idx_user_comment,uniqueIndex" json:"comment_id"`
}

