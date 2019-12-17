package badger

import "github.com/dgraph-io/badger"

type bgStore struct {
	db *badger.DB
}
