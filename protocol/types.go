package protocol

import (
	"net"
	"net/textproto"
)

type DB interface {
	Close() error

	// kv
	Del([]byte) error
	Set([]byte, []byte) error
	Get([]byte) ([]byte, error)

	// list
	Lpush([]byte, []byte) (uint64, error)
	Lrange([]byte, uint64, uint64) ([][]byte, error)
}

type Server interface {
	Run()
	Stop()
}

type responseWriter interface {
	respError(error)
	respInteger(int)
	respBulk(string)
	respSimple(string)
	respArray([][]byte)
}

type server struct {
	db  DB
	lis net.Listener
	ch  chan struct{}
}

type request struct {
	args [][]byte
}

type response struct {
	w *textproto.Writer
}

type dealFunc (func(DB, responseWriter, [][]byte) error)

var dealRegister map[string]dealFunc
