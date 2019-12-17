package bolt

import (
	"github.com/boltdb/bolt"
	"github.com/infinivision/gaeadbBench/protocol"
)

func New(name string) protocol.DB {
	if db, err := bolt.Open(name, 0600, &bolt.Options{NoGrowSync: false}); err != nil {
		return nil
	} else {
		if err := db.Update(func(tx *bolt.Tx) error {
			if _, err := tx.CreateBucketIfNotExists([]byte("test")); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return nil
		}
		return &btStore{db}
	}
}

func (db *btStore) Close() error {
	return db.db.Close()
}

func (db *btStore) Del(k []byte) error {
	return db.db.Update(func(tx *bolt.Tx) error { return tx.Bucket([]byte("test")).Delete(k) })
}

func (db *btStore) Set(k, v []byte) error {
	return db.db.Update(func(tx *bolt.Tx) error { return tx.Bucket([]byte("test")).Put(k, v) })
}

func (db *btStore) Get(k []byte) ([]byte, error) {
	var v []byte

	if err := db.db.View(func(tx *bolt.Tx) error { v = tx.Bucket([]byte("test")).Get(k); return nil }); err != nil {
		return nil, err
	}
	return v, nil
}
