package badger

import (
	"encoding/binary"
	"sync/atomic"

	"github.com/dgraph-io/badger"
	"github.com/infinivision/gaeadbBench/protocol"
)

func New(name string) protocol.DB {
	opts := badger.DefaultOptions
	opts.SyncWrites = true
	opts.Dir = name
	opts.ValueDir = name
	if db, err := badger.Open(opts); err != nil {
		return nil
	} else {
		return &bgStore{db: db}
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

func (db *bgStore) Lpush(k, v []byte) (uint64, error) {
	tx := db.db.NewTransaction(true)
	defer tx.Discard()
	if n, err := lpush(tx, k, v, atomic.AddUint64(&db.cnt, 1)); err != nil {
		return 0, err
	} else {
		if err = tx.Commit(); err != nil {
			return 0, err
		}
		return n, nil
	}
}

func (db *bgStore) Lrange(k []byte, start, end uint64) ([][]byte, error) {
	tx := db.db.NewTransaction(false)
	defer tx.Discard()
	return lrange(tx, k, start, end)
}

func lpush(tx *badger.Txn, k, v []byte, idx uint64) (uint64, error) {
	if err := tx.Set(eListKey(k, idx), v); err != nil {
		return 0, err
	}
	return idx, nil
}

func lrange(tx *badger.Txn, k []byte, start, end uint64) ([][]byte, error) {
	var vs [][]byte

	iopt := badger.IteratorOptions{}
	iopt.Prefix = eListMetaKey(k)
	iopt.PrefetchValues = true
	itr := tx.NewIterator(iopt)
	defer itr.Close()
	for itr.Seek(eListMetaKey(k)); itr.Valid(); itr.Next() {
		k := itr.Item().KeyCopy(nil)
		idx := dListMetaValue(k[len(k)-8:])
		if idx < start || idx > end {
			break
		}
		if idx >= start && idx <= end {
			v, err := itr.Item().ValueCopy(nil)
			if err != nil {
				return nil, err
			}
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
