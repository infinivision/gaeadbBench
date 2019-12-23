package gaea

import (
	"encoding/binary"
	"sync/atomic"

	"github.com/infinivision/gaeadb/db"
	"github.com/infinivision/gaeadb/errmsg"
	"github.com/infinivision/gaeadb/transaction"
	"github.com/infinivision/gaeadbBench/protocol"
)

func New(name string) protocol.DB {
	cfg := db.DefaultConfig()
	cfg.DirName = name
	cfg.CacheSize = 60000
	if db, err := db.Open(cfg); err != nil {
		return nil
	} else {
		return &gaStore{db: db}
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

func (db *gaStore) Lpush(k, v []byte) (uint64, error) {
	tx, err := db.db.NewTransaction(false)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	if n, err := lpush(tx, k, v, atomic.AddUint64(&db.cnt, 1)); err != nil {
		return 0, err
	} else {
		if err = tx.Commit(); err != nil {
			return 0, err
		}
		return n, nil
	}
}

func (db *gaStore) Lrange(k []byte, start, end uint64) ([][]byte, error) {
	tx, err := db.db.NewTransaction(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	return lrange(tx, k, start, end)
}

func lpush(tx transaction.Transaction, k, v []byte, idx uint64) (uint64, error) {
	if err := tx.Set(eListKey(k, idx), v); err != nil {
		return 0, err
	}
	return idx, nil
}

func lrange(tx transaction.Transaction, k []byte, start, end uint64) ([][]byte, error) {
	var vs [][]byte

	itr, err := tx.NewForwardIterator(eListMetaKey(k))
	if err != nil {
		return nil, err
	}
	defer itr.Close()
	for itr.Valid() {
		k := itr.Key()
		idx := dListMetaValue(k[len(k)-8:])
		if idx < start || idx > end {
			break
		}
		if idx >= start && idx <= end {
			v, err := itr.Value()
			if err != nil {
				return nil, err
			}
			vs = append(vs, v)
		}
		if err := itr.Next(); err != nil {
			if err == errmsg.ScanEnd {
				break
			}
			return nil, err
		}

	}
	return vs, nil
}

// 'l' + k
func eListMetaKey(k []byte) []byte {
	return append([]byte{'l'}, k...)
}

func dListMetaKey(buf []byte) []byte {
	return buf[1:]
}

// index
func eListMetaValue(idx uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, idx)
	return buf
}

func dListMetaValue(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}

// 'l' + k + index
func eListKey(k []byte, idx uint64) []byte {
	buf := []byte{}
	buf = append([]byte{'l'}, k...)
	buf = append(buf, eListMetaValue(idx)...)
	return buf
}

func dListKey(buf []byte) []byte {
	n := len(buf)
	return buf[1 : n-8]
}

func dListIndex(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}
