package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content string `json:"content" gorm:"type:text; not null; default:null"`
	Likes   int    `json:"likes" gorm:"type:int; not null; default:0"`
	UserID  uint   `json:"user_id" gorm:"type:int;not null; default:1"`
}

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"type:text; not null; default:null:"`
	Password  string `json:"password" gorm:"type:text; not null; default:null"`
	Followers int    `json:"followers" gorm:"type:int; not null; default:0"`
}

type Comment struct {
	gorm.Model
	Content string `json:"content" gorm:"text; not null; default:null"`
	Likes   int    `json:"likes" gorm:"type:int;not null; default:0"`
	Author  int    `json:"author" gorm:"type:int; not null; default:1"`
	Post_id int    `json:"post_id" gorm:"type:int; not null; default:1"`
}
