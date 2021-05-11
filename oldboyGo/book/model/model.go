package model

type Book struct {
	Id int64 `db:"id"`
	Title string `db:"title"`
	Price int64 `db:"price"`
}

