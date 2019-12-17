package badger

import (
	"github.com/dgraph-io/badger"
	"github.com/infinivision/gaeadbBench/protocol"
)

func New(name string) protocol.DB {
	opts := badger.DefaultOptions(name)
	opts.SyncWrites = true
	if db, err := badger.Open(opts); err != nil {
		return nil
	} else {
		return &bgStore{db}
	}
}

func (db *bgStore) Close() error {
	return db.db.Close()
}

func (db *bgStore) Del(k []byte) error {
	tx := db.db.NewTransaction(true)
	defer tx.Discard()
	if err := tx.Delete(k); err != nil {
		return err
	}
	return tx.Commit()
}

func (db *bgStore) Set(k, v []byte) error {
	tx := db.db.NewTransaction(true)
	defer tx.Discard()
	if err := tx.Set(k, v); err != nil {
		return err
	}
	return tx.Commit()
}

func (db *bgStore) Get(k []byte) ([]byte, error) {
	tx := db.db.NewTransaction(false)
	defer tx.Discard()
	if it, err := tx.Get(k); err != nil {
		return nil, err
	} else {
		return it.ValueCopy(nil)
	}
}
