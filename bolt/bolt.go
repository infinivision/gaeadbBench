package bolt

import (
	"bytes"
	"encoding/binary"
	"sync/atomic"

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
		return &btStore{db: db}
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

func (db *btStore) Lpush(k, v []byte) (uint64, error) {
	tx, err := db.db.Begin(true)
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

func (db *btStore) Lrange(k []byte, start, end uint64) ([][]byte, error) {
	tx, err := db.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	return lrange(tx, k, start, end)
}

func lpush(tx *bolt.Tx, k, v []byte, idx uint64) (uint64, error) {
	if err := tx.Bucket([]byte("test")).Put(eListKey(k, idx), v); err != nil {
		return 0, err
	}
	return idx, nil
}

func lrange(tx *bolt.Tx, k []byte, start, end uint64) ([][]byte, error) {
	var vs [][]byte

	itr := tx.Bucket([]byte("test")).Cursor()
	prefix := eListMetaKey(k)
	for k, v := itr.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = itr.Next() {
		k = append([]byte{}, k...)
		idx := dListMetaValue(k[len(k)-8:])
		if idx < start || idx > end {
			break
		}
		if idx >= start && idx <= end {
			v = append([]byte{}, v...)
			vs = append(vs, v)
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
