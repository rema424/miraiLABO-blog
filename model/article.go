package model

import "time"

type Article struct {
	//ID    int    `db:"id"`
	//Title string `db:"title"`
	ID      int       `db:"id" form:"id"`
	Title   string    `db:"title" form:"title"`
	Body    string    `db:"body" form:"body"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}
