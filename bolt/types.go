package bolt

import "github.com/boltdb/bolt"

type btStore struct {
	db *bolt.DB
}
