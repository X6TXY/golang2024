package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Followers int       `json:"followers"`
	Posts     []Post    `json:"posts"`
	Comments  []Comment `json:"comments"`
}

type Post struct {
	gorm.Model
	Content  string    `json:"content"`
	Likes    int       `json:"likes"`
	UserID   uint      `json:"user_id"`
	User     User      `json:"user" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	Likes   int    `json:"likes"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`
	Post    Post   `json:"post" gorm:"foreignKey:PostID"`
}
