package types

import (
	"time"

	_ "github.com/lib/pq"
)

type Post struct {
	PostKey     string `db:"post_key"`
	Title       string `db:"title"`
	Content     string
	CreatedDate time.Time `db:"created_date"`
	UpdatedDate time.Time `db:"updated_date"`
}
