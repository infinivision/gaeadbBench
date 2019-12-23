package gaea

import "github.com/infinivision/gaeadb/db"

type gaStore struct {
	cnt uint64
	db  db.DB
}
