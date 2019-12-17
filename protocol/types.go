package protocol

import (
	"net"
	"net/textproto"

	"github.com/nnsgmsone/units/breaker"
)

type DB interface {
	Close() error

	Del([]byte) error
	Set([]byte, []byte) error
	Get([]byte) ([]byte, error)
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
	brk breaker.Breaker
}

type connection struct {
	state int
	s     *server
	conn  net.Conn
}

type request struct {
	args [][]byte
}

type response struct {
	w *textproto.Writer
}

type dealFunc (func(DB, responseWriter, [][]byte))

var dealRegister map[string]dealFunc
