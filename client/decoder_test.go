package client

import (
	"testing"
)

func TestDecodeExtractInt(t *testing.T) {
	b := []byte{123, 0, 0, 0}
	v, err := extractInt(&b)
	if err != nil {
		t.Fatal("Extracting of int failed")
	}

	if v != 123 {
		t.Fatal("Extracting of int returned wrong value")
	}

	if len(b) != 0 {
		t.Fatal("Extracting of int returned non-empty slice")
	}
}

func TestDecodeExtractInt64(t *testing.T) {
	b := []byte{123, 0, 0, 0, 0, 0, 0, 0}
	v, err := extractInt64(&b)
	if err != nil {
		t.Fatal("Extracting of int failed")
	}

	if v != 123 {
		t.Fatal("Extracting of int returned wrong value")
	}

	if len(b) != 0 {
		t.Fatal("Extracting of int returned non-empty slice")
	}
}

func TestDecoderExtractString(t *testing.T) {
	b := []byte{
		11, 0, 0, 0,
		116, 101, 115, 116, 32, 115, 116, 114, 105, 110, 103,
	}
	v, err := extractString(&b)
	if err != nil {
		t.Fatal("Extracting of string failed")
	}

	if v != "test string" {
		t.Fatal("Extracting of string returned wrong value")
	}
}

func TestDecoderDecodeHeader(t *testing.T) {
	b := []byte{
		11, 0, 0, 0, // svc_id
		22, 0, 0, 0, // body_length
		33, 130, 101, 77, // requestId
	}

	header, err := decodeHeader(&b)
	if err != nil {
		t.Fatal("Extracting of string failed")
	}

	if header.svcId != 11 {
		t.Fatal("Extracting of header returned wrong svc_id value")
	}

	if header.bodyLength != 22 {
		t.Fatal("Extracting of header returned wrong body_length value")
	}

	if header.requestId != 1298498081 {
		t.Fatal("Extracting of header returned wrong request_id value")
	}
}

func TestDecoderDecodeNilResponse(t *testing.T) {
	_, err := decodeBody(nil)
	if err == nil {
		t.Fatal("Extracting of nil response succeed")
	}
}

func TestDecoderDecodeBrokenResponseBody(t *testing.T) {
	b := []byte{
		0, 0, 0, // broken return code
	}

	_, err := decodeBody(&b)
	if err == nil {
		t.Fatal("Extracting of broken response body succeed")
	}
}

func TestDecoderDecodeBrokenResponseBody2(t *testing.T) {
	b := []byte{
		0, 0, 0, 0, // return code == 0
		17, 0, 0, // broken string length
	}

	_, err := decodeBody(&b)
	if err == nil {
		t.Fatal("Extracting of broken response body succeed")
	}
}

func TestDecoderDecodeBrokenResponseBody3(t *testing.T) {
	b := []byte{
		0, 0, 0, 0, // return code == 0
		17, 0, 0, 0,
		99, 108, 105, 101, 110, // broken string
	}

	_, err := decodeBody(&b)
	if err == nil {
		t.Fatal("Extracting of broken response body succeed")
	}
}

func TestDecoderDecodeBrokenResponseBody4(t *testing.T) {
	b := []byte{
		0, 0, 0, 0, // return code == 0
		17, 0, 0, 0,
		99, 108, 105, 101, 110,
		116, 32, 105, 100, 101, 110,
		116, 105, 102, 105, 101, 114, // client_id == "client identifier"
		4, 208, 34, 0, // client_type == 2281476
		7, 0, 0, 0,
		97, 119, 101, 115, 111, 109, 101, // username == "awesome"
		192, 146, 80, 97, // expires_in == 1632670400
		32, 83, 105, 0, 13, // broken user_id
	}

	_, err := decodeBody(&b)
	if err == nil {
		t.Fatal("Extracting of broken response body succeed")
	}
}

func TestDecoderDecodeResponseBody(t *testing.T) {
	b := []byte{
		0, 0, 0, 0, // return code == 0
		17, 0, 0, 0,
		99, 108, 105, 101, 110,
		116, 32, 105, 100, 101, 110,
		116, 105, 102, 105, 101, 114, // client_id == "client identifier"
		4, 208, 34, 0, // client_type == 2281476
		7, 0, 0, 0,
		97, 119, 101, 115, 111, 109, 101, // username == "awesome"
		192, 146, 80, 97, // expires_in == 1632670400
		32, 83, 105, 0, 13, 1, 58, 0, // user_id == 16326704002323232
	}

	body, err := decodeBody(&b)
	if err != nil {
		t.Fatal("Extracting of string failed")
	}

	expected := CubeResponseBody{
		ReturnCode: 0,
		ClientId:   "client identifier",
		ClientType: 2281476,
		Username:   "awesome",
		ExpiresIn:  1632670400,
		UserId:     16326704002323232,
	}

	if *body != expected {
		t.Fatal("Extracting of body returned wrong field(-s) value(-es)")
	}
}

func TestDecoderDecodeErrorResponseBody(t *testing.T) {
	b := []byte{
		1, 0, 0, 0, // return code == 1
		17, 0, 0, 0,
		99, 108, 105, 101, 110,
		116, 32, 105, 100, 101, 110,
		116, 105, 102, 105, 101, 114, // error_msg == "client identifier"
	}

	body, err := decodeBody(&b)
	if err != nil {
		t.Fatal("Extracting of string failed")
	}

	expected := CubeResponseBody{
		ReturnCode: 1,
		ErrString:  "client identifier",
	}

	if *body != expected {
		t.Fatal("Extracting of body returned wrong field(-s) value(-es)")
	}
}
