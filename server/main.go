package main

import (
	"log"
	"runtime"
	"syscall"

	"net/http"
	_ "net/http/pprof"

	"github.com/infinivision/gaeadbBench/gaea"
	"github.com/infinivision/gaeadbBench/protocol"
)

func main() {
	{
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}
	enlargelimit()
	s := protocol.New(6378, gaea.New("test.db"))
	//s := protocol.New(6378, bolt.New("test.db"))
	//s := protocol.New(6378, badger.New("test.db"))
	s.Run()
}

func enlargelimit() error {
	var rlimit syscall.Rlimit

	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		return err
	} else {
		rlimit.Cur = rlimit.Max
		return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	}
	return nil
}
