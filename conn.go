package riakpbc

import (
	"net"
)

type Conn struct {
	addr string
	conn net.Conn
}

// Dial connects to a single riak server.
func Dial(addr string) (conns []Conn, err error) {
  num := 5

  for i := 0; i < num; i++ {
    var c Conn
    c.addr = addr
    c.conn, err = net.Dial("tcp", addr)

    _ = append(conns, c)

    if err != nil {
      return nil, err
    }
  }

	return conns, nil
}

func (c *Conn) Close() {
	c.conn.Close()
}
