package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestEncoderEncodeInt(t *testing.T) {
	b := make([]byte, 0)
	err := appendInt(&b, 10)
	if err != nil {
		t.Fatal("Encoding of int failed")
	}

	expected := []byte{10, 0, 0, 0}
	if bytes.Compare(b, expected) != 0 {
		t.Fatal("Encoding of int failed")
	}
}

func TestEncoderEncodeString(t *testing.T) {
	b := make([]byte, 0)
	err := appendString(&b, "test string")
	if err != nil {
		t.Fatal("Encoding of int failed")
	}

	expected := []byte{11, 0, 0, 0,
		116, 101, 115, 116, 32, 115, 116, 114, 105, 110, 103,
	}
	if bytes.Compare(b, expected) != 0 {
		t.Fatal("Encoding of int failed")
	}
}

func TestEncoderEmptyReq(t *testing.T) {
	e := Encoder{}
	_, _, err := e.FormatRequest(nil)

	if err == nil {
		t.Fatal("Encoding of empty request successed")
	}

	if fmt.Sprint(err) != "empty ptr for request passed to encoder" {
		t.Fatal("Encoding of empty request failed with wrong msg")
	}
}

func TestEncoderStdReq(t *testing.T) {
	request := CubeRequestBody{}
	e := Encoder{}
	_, _, err := e.FormatRequest(&request)
	if err != nil {
		t.Fatal("Encoding of empty request failed")
	}
}

func TestEncoderReq(t *testing.T) {
	request := CubeRequestBody{Token: "token", Scope: "scope", SvcId: 0}
	e := Encoder{}
	reqId, reqBin, _ := e.FormatRequest(&request)

	expected := []byte{
		0, 0, 0, 0, // svc_id
		22, 0, 0, 0, // body length
		0, 0, 0, 0, // req_id
		1, 0, 0, 0, // svc_msg
		5, 0, 0, 0, 116, 111, 107, 101, 110, // string "token"
		5, 0, 0, 0, 115, 99, 111, 112, 101, // string "scope"
	}
	binary.LittleEndian.PutUint32(expected[8:], uint32(reqId))

	if bytes.Compare(reqBin, expected) != 0 {
		t.Fatalf("Encoding of request %+v failed", request)
	}
}
