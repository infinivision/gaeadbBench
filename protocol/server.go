package protocol

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
)

func New(port int, db DB) Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil
	}
	return &server{db, lis, make(chan struct{})}
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
			go s.process(conn)
		}
	}
}

func (s *server) Stop() {
	s.ch <- struct{}{}
	<-s.ch
	close(s.ch)
	return
}

func (s *server) process(c net.Conn) {
	for {
		resp := &response{textproto.NewWriter(bufio.NewWriter(c))}
		req, err := readRequest(bufio.NewReader(c))
		if err != nil {
			return
		}
		f, ok := dealRegister[string(req.args[0])]
		if !ok {
			err = fmt.Errorf("unknown command '%s'", string(req.args[0]))
			resp.respError(err)
			return
		}
		if err := f(s.db, resp, req.args[1:]); err != nil {
			//		fmt.Printf("failed to deal %s\n", string(req.args[0]))
		}
	}
}
