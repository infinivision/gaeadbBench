package bolt

import "github.com/boltdb/bolt"

type btStore struct {
	cnt uint64
	db  *bolt.DB
}
