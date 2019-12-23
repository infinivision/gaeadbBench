package badger

import "github.com/dgraph-io/badger"

type bgStore struct {
	cnt uint64
	db  *badger.DB
}
