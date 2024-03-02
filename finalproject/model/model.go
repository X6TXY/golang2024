package model

import "time"

type Post struct {
	ID          int       `json:"id"`
	Text        string    `json:"text"`
	Date        time.Time `json:"date"`
	Update_date time.Time `json:"update_date"`
	//Likes int `json:"likes"`
	//Author_id int `json:"author_id"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Date      time.Time `json:"date"`
	Followers int       `json:"followers"`
}

type Comment struct {
	ID          int       `json:"id"`
	Post_id     int       `json:"post_id"`
	Date        time.Time `json:"date"`
	Content     string    `json:"content"`
	Likes       int       `json:"likes"`
	User_id     int       `json:"user_id"`
}
