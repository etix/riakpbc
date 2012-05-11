package riakpbc

import (
	"errors"
)

var (
	ErrLengthZero    = errors.New("length response 0")
	ErrCorruptHeader = errors.New("corrupt header")
)

func writeRequest(c *Conn, formattedRequest []byte) (err error) {
	_, err = c.conn.Write(formattedRequest)
	return err
}

func readResponse(c *Conn) (respraw []byte, err error) {
	respraw = make([]byte, 512)

	c.conn.Read(respraw)

	_ = respraw[3]

	return respraw, nil
}
