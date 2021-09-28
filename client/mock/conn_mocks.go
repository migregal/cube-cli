package mock

import (
	"errors"
	"io"
	"net"
	"time"
)

type Conn struct {
	WriteError bool
	Writtable  bool
	ReadError  bool
	Response   []byte
}

func (c Conn) Read(b []byte) (n int, err error) {
	if c.ReadError {
		return 0, errors.New("unable to read")
	}
	copy(b, c.Response)
	return len(c.Response), io.EOF
}

func (c Conn) Write(b []byte) (n int, err error) {
	if c.WriteError {
		return 0, errors.New("unable to write")
	}
	if c.Writtable {
		return len(b), nil
	}
	return 0, nil
}

func (c Conn) Close() error {
	return nil
}

func (c Conn) LocalAddr() net.Addr {
	return nil
}

func (c Conn) RemoteAddr() net.Addr {
	return nil
}

func (c Conn) SetDeadline(t time.Time) error {
	return nil
}

func (c Conn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c Conn) SetWriteDeadline(t time.Time) error {
	return nil
}
