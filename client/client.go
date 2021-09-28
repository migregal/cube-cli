package client

import (
	"bytes"
	"errors"
	"io"
	"net"
	"time"
)

const TokenVerifySvcMsg = 0x00000001

type cubeClient struct {
	svcId Int
	conn  net.Conn
}

func NewConnection(svcId Int, host string, port string) (*cubeClient, error) {
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

	sReqId, err := c.sendRequest(token, scope)
	if err != nil {
		return nil, err
	}

	reqId, resp, err := c.readResponse()
	if err != nil {
		return nil, err
	}

	if reqId != sReqId {
		return nil, errors.New("received response for other request")
	}
	return resp, nil
}

func (c *cubeClient) sendRequest(token, scope string) (reqId RId, err error) {
	req := CubeRequestBody{SvcId: c.svcId, Token: token, Scope: scope}
	reqId, bin, err := Encoder{}.FormatRequest(&req)
	if err != nil {
		return
	}

	n, err := c.conn.Write(bin)
	if err != nil {
		return
	}

	if n != len(bin) {
		err = errors.New("failed to send request")
		return
	}

	return
}

func (c *cubeClient) readResponse() (reqId RId, resp *CubeResponseBody, err error) {
	var buf bytes.Buffer
	if _, err = io.Copy(&buf, c.conn); err != nil {
		return
	}

	reqId, resp, err = Decoder{}.DecodeResponse(buf.Bytes())
	if err != nil {
		return 0, nil, errors.New("failed to decode response")
	}

	return
}

func (c *cubeClient) CloseConnection() error {
	return c.conn.Close()
}
