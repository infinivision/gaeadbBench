package gaea

import (
	"github.com/infinivision/gaeadb/db"
	"github.com/infinivision/gaeadbBench/protocol"
)

func New(name string) protocol.DB {
	cfg := db.DefaultConfig()
	cfg.DirName = name
	cfg.CacheSize = 65536
	if db, err := db.Open(cfg); err != nil {
		return nil
	} else {
		return &gaStore{db}
	}
}

func (db *gaStore) Close() error {
	return db.db.Close()
}

func (db *gaStore) Del(k []byte) error {
	return db.db.Del(k)
}

func (db *gaStore) Set(k, v []byte) error {
	return db.db.Set(k, v)
}

func (db *gaStore) Get(k []byte) ([]byte, error) {
	return db.db.Get(k)
}
