package protocol

import (
	"errors"
	"strconv"
)

func init() {
	dealRegister = make(map[string]dealFunc)

	dealRegister["SET"] = deal0
	dealRegister["DEL"] = deal1
	dealRegister["GET"] = deal2

	dealRegister["LPUSH"] = deal3
	dealRegister["LRANGE"] = deal4
}

// SET k v
func deal0(db DB, resp responseWriter, args [][]byte) error {
	if len(args) < 2 {
		resp.respError(errors.New("wrong number of arguments for 'SET' command"))
		return errors.New("wrong number of arguments for 'SET' command")
	}
	if err := db.Set(args[0], args[1]); err != nil {
		resp.respError(err)
		return err
	}
	resp.respSimple("OK")
	return nil
}

// DEL k
func deal1(db DB, resp responseWriter, args [][]byte) error {
	if len(args) < 1 {
		resp.respError(errors.New("wrong number of arguments for 'DEL' command"))
		return errors.New("wrong number of arguments for 'DEL' command")
	}
	if err := db.Del(args[0]); err != nil {
		resp.respInteger(0)
		return err
	} else {
		resp.respInteger(1)
		return nil
	}
}

// GET k
func deal2(db DB, resp responseWriter, args [][]byte) error {
	if len(args) < 1 {
		resp.respError(errors.New("wrong number of arguments for 'GET' command"))
		return errors.New("wrong number of arguments for 'GET' command")
	}
	v, err := db.Get(args[0])
	switch {
	case err != nil:
		resp.respBulk("")
		return err
	default:
		resp.respBulk(string(v))
		return nil
	}
}

// LPUSH k v
func deal3(db DB, resp responseWriter, args [][]byte) error {
	if len(args) < 2 {
		resp.respError(errors.New("wrong number of arguments for 'LPUSH' command"))
		return errors.New("wrong number of arguments for 'LPUSH' command")
	}
	if n, err := db.Lpush(args[0], args[1]); err != nil {
		resp.respError(err)
		return err
	} else {
		resp.respInteger(int(n))
		return nil
	}
}

// LRANGE k start stop
func deal4(db DB, resp responseWriter, args [][]byte) error {
	if len(args) < 3 {
		resp.respError(errors.New("wrong number of arguments for 'LRANE' command"))
		return errors.New("wrong number of arguments for 'LRANE' command")
	}
	start, err := strconv.ParseInt(string(args[1]), 0, 64)
	if err != nil {
		resp.respError(err)
		return err
	}
	stop, err := strconv.ParseInt(string(args[2]), 0, 64)
	if err != nil {
		resp.respError(err)
		return err
	}
	vs, err := db.Lrange(args[0], uint64(start), uint64(stop))
	switch {
	case err == nil:
		resp.respArray(vs)
		return nil
	default:
		resp.respArray([][]byte{})
		return err
	}
}
