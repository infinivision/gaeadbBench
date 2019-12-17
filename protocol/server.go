package protocol

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"net/textproto"

	"github.com/nnsgmsone/units/breaker"
)

func New(port int, db DB) Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil
	}
	cfg := breaker.DefaultConfig()
	cfg.MaxRequests = 1000000 // 100w tps / 10s
	return &server{db, lis, make(chan struct{}), breaker.New(cfg)}
}

func (s *server) Run() {
	ch := make(chan net.Conn)
	go func() {
		for {
			conn, err := s.lis.Accept()
			if err != nil {
				for {
					lis, err := net.Listen("tcp", s.lis.Addr().String())
					if err == nil {
						s.lis.Close()
						s.lis = lis
						break
					}
				}
			}
			ch <- conn
		}
	}()
	for {
		select {
		case <-s.ch:
			s.ch <- struct{}{}
			return
		case conn := <-ch:
			go s.brk.NewConnection(&connection{0, s, conn})
		}
	}
}

func (s *server) Stop() {
	s.ch <- struct{}{}
	<-s.ch
	close(s.ch)
	return
}

func (c *connection) Response() error {
	resp := &response{textproto.NewWriter(bufio.NewWriter(c.conn))}
	resp.respSimple("ok")
	req, err := readRequest(bufio.NewReader(c.conn))
	if err != nil {
		c.state = -1
		if err != io.EOF {
			resp.respError(errors.New("illegal request"))
		}
		return err
	}
	f, ok := dealRegister[string(req.args[0])]
	if !ok {
		c.state = -1
		err = fmt.Errorf("unknown command '%s'", string(req.args[0]))
		resp.respError(err)
		return err
	}
	f(c.s.db, resp, req.args[1:])
	return nil
}

func (c *connection) Close() error {
	return c.conn.Close()
}

func (c *connection) Serve() error {
	for {
		if err := c.s.brk.NewRequest(c); err != nil {
			resp := &response{textproto.NewWriter(bufio.NewWriter(c.conn))}
			resp.respError(err)
			break
		}
		if c.state != 0 {
			break
		}
	}
	return nil
}
