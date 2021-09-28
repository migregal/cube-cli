package client

import (
	"testing"
)

func TestUnableToSendRequest(t *testing.T) {
	c := cubeClient{svcId: 0, conn: EmptyWriteMockConn{}}
	defer c.CloseConnection()

	_, err := c.VerifyToken("token", "scope")
	if err == nil {
		t.Fatal("Failed writing passed")
	}
	if err.Error() != "failed to send request" {
		t.Fatal("Wrong error msg")
	}
}

func TestUnableToReadResponse(t *testing.T) {
	c := cubeClient{svcId: 0, conn: EmptyReadMockConn{}}
	defer c.CloseConnection()

	_, err := c.VerifyToken("token", "scope")
	if err == nil {
		t.Fatal("Failed reading passed")
	}
	if err.Error() != "unable to read" {
		t.Fatal("Wrong error msg")
	}
}

func TestWrongResponseHeaderFormat(t *testing.T) {
	c := cubeClient{svcId: 0, conn: BrokenRespFmtMockConn{respLen: 10}}
	defer c.CloseConnection()

	_, err := c.VerifyToken("token", "scope")
	if err == nil {
		t.Fatal("Failed decoding passed")
	}
	if err.Error() != "failed to decode response" {
		t.Fatal("Wrong error msg")
	}
}

func TestWrongResponseBodyFormat(t *testing.T) {
	c := cubeClient{svcId: 0, conn: BrokenRespFmtMockConn{respLen: 13}}
	defer c.CloseConnection()

	_, err := c.VerifyToken("token", "scope")
	if err == nil {
		t.Fatal("Failed decoding passed")
	}
	if err.Error() != "failed to decode response" {
		t.Fatal("Wrong error msg")
	}
}

func TestWrongReqIdResponseHandling(t *testing.T) {
	c := cubeClient{svcId: 0, conn: FixedResponseMockConn{expected: []byte{
		0, 0, 0, 0,
		25, 0, 0, 0,
		0, 0, 0, 0, // header
		1, 0, 0, 0, // return code == 1
		17, 0, 0, 0,
		99, 108, 105, 101, 110,
		116, 32, 105, 100, 101, 110,
		116, 105, 102, 105, 101, 114, // error_msg == "client identifier"
	}}}
	defer c.CloseConnection()

	_, err := c.VerifyToken("token", "scope")
	if err == nil {
		t.Fatal("Wrong reqId response passed")
	}
	if err.Error() != "received response for other request" {
		t.Fatal("Wrong error message")
	}
}
