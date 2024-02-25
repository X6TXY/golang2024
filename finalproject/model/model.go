//Model sructure of the project 

package model

import "time"

type Comment struct {
	ID   int       `json:"id"`
	Text string    `json:"text"`
	Date time.Time `json:"date"`
}

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password []byte `json:"password"` 
}