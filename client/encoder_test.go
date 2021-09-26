package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestEncoderEmptyReq(t *testing.T) {
	e := Encoder{}
	_, _, err := e.FormatRequest(nil)

	if err == nil {
		t.Fatalf("Encoding of empty request successed")
	}

	if fmt.Sprint(err) != "empty ptr for request passed to encoder" {
		t.Fatalf("Encoding of empty request failed with wrong msg")
	}
}

func TestEncoderStdReq(t *testing.T) {
	request := CubeRequest{}
	e := Encoder{}
	_, _, err := e.FormatRequest(&request)
	if err != nil {
		t.Fatalf("Encoding of empty request failed")
	}
}

func TestEncoderReq(t *testing.T) {
	request := CubeRequest{Token: "token", Scope: "scope", SvcId: 0}
	e := Encoder{}
	reqId, reqBin, _ := e.FormatRequest(&request)

	reqIdB := make([]byte, 4)
	binary.LittleEndian.PutUint32(reqIdB, uint32(reqId))
	expected := []byte{
		0, 0, 0, 0, // svc_id
		22, 0, 0, 0, // body length
		reqIdB[0], reqIdB[1], reqIdB[2], reqIdB[3], // req_id
		1, 0, 0, 0, // svc_msg
		5, 0, 0, 0, 116, 111, 107, 101, 110, // string "token"
		5, 0, 0, 0, 115, 99, 111, 112, 101, // string "scope"
	}

	if bytes.Compare(reqBin, expected) != 0 {
		t.Fatalf("Encoding of request %+v failed", request)
	}
}
