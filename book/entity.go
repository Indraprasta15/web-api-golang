package book

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	Discount    int
	CreatedAt   time.Time
	Updated     time.Time
}
