package client

import (
	"bytes"
	"errors"
	"io"
	"net"
)

const TokenVerifySvcMsg = 0x00000001

type cubeClient struct {
	svcId int32
	addr  *net.TCPAddr
}

func NewClient(svcId int32, host string, port string) (*cubeClient, error) {
	addr, err := net.ResolveTCPAddr("tcp", host+":"+port)
	if err != nil {
		return nil, err
	}

	client := cubeClient{svcId: svcId, addr: addr}
	return &client, nil
}

func (c *cubeClient) VerifyToken(token, scope string) (*CubeResponseBody, error) {
	conn, err := net.DialTCP("tcp", nil, c.addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	req := CubeRequestBody{SvcId: c.svcId, Token: token, Scope: scope}
	_, bin, err := Encoder{}.FormatRequest(&req)
	if err != nil {
		return nil, err
	}

	n, err := conn.Write(bin)
	if err != nil {
		return nil, err
	}

	if n != len(bin) {
		return nil, errors.New("failed to send request")
	}

	var buf bytes.Buffer
	if _, err = io.Copy(&buf, conn); err != nil {
		return nil, err
	}

	var resp *CubeResponseBody
	resp, err = Decoder{}.DecodeResponse(buf.Bytes())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
