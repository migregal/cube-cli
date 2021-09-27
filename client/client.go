package client

import (
	"bytes"
	"errors"
	"io"
	"net"
	"time"
)

const TokenVerifySvcMsg = 0x00000001

type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
	SetDeadline(t time.Time) error
}

type cubeClient struct {
	svcId int32
	conn  Conn
}

func NewConnection(svcId int32, host string, port string) (*cubeClient, error) {
	addr, err := net.ResolveTCPAddr("tcp", host+":"+port)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}

	client := cubeClient{svcId: svcId, conn: conn}
	return &client, nil
}

func (c *cubeClient) VerifyToken(token, scope string) (*CubeResponseBody, error) {
	err := c.conn.SetDeadline(time.Now().Add(time.Second * 5))
	if err != nil {
		return nil, err
	}

	req := CubeRequestBody{SvcId: c.svcId, Token: token, Scope: scope}
	sReqId, bin, err := Encoder{}.FormatRequest(&req)
	if err != nil {
		return nil, err
	}

	n, err := c.conn.Write(bin)
	if err != nil {
		return nil, err
	}

	if n != len(bin) {
		return nil, errors.New("failed to send request")
	}

	var buf bytes.Buffer
	if _, err = io.Copy(&buf, c.conn); err != nil {
		return nil, err
	}

	var resp *CubeResponseBody
	reqId, resp, err := Decoder{}.DecodeResponse(buf.Bytes())
	if err != nil {
		return nil, errors.New("failed to decode response")
	}

	if reqId != sReqId {
		return nil, errors.New("received response for other request")
	}
	return resp, nil
}

func (c *cubeClient) CloseConnection() error {
	return c.conn.Close()
}
