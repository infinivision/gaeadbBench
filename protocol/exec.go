package protocol

import (
	"errors"
)

func init() {
	dealRegister = make(map[string]dealFunc)

	dealRegister["SET"] = deal0
	dealRegister["DEL"] = deal1
	dealRegister["GET"] = deal2
}

// SET k v
func deal0(db DB, resp responseWriter, args [][]byte) {
	if len(args) < 2 {
		resp.respError(errors.New("wrong number of arguments for 'SET' command"))
		return
	}
	if err := db.Set(args[0], args[1]); err != nil {
		resp.respError(err)
		return
	}
	resp.respSimple("OK")
}

// DEL k
func deal1(db DB, resp responseWriter, args [][]byte) {
	if len(args) < 1 {
		resp.respError(errors.New("wrong number of arguments for 'DEL' command"))
		return
	}
	if err := db.Del(args[0]); err != nil {
		resp.respInteger(0)
	} else {
		resp.respInteger(1)
	}
}

// GET k
func deal2(db DB, resp responseWriter, args [][]byte) {
	if len(args) < 1 {
		resp.respError(errors.New("wrong number of arguments for 'GET' command"))
		return
	}
	v, err := db.Get(args[0])
	switch {
	case err != nil:
		resp.respBulk("")
	default:
		resp.respBulk(string(v))
	}
}
