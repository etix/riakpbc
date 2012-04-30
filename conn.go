package riakpbc

import (
	"net"
)

type Conn struct {
	addr string
	conn net.Conn
}

// Dial connects to a single riak server.
func Dial(addr string) (*Conn, error) {
	var c Conn
	var err error
	c.addr = addr
	c.conn, err = net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Conn) Close() {
	c.conn.Close()
}
