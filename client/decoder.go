package client

import (
	"encoding/binary"
	"errors"
)

type Decoder struct{}

func (d Decoder) DecodeResponse(response []byte) (RId, *CubeResponseBody, error) {
	b := make([]byte, len(response))
	if copy(b, response) != len(response) {
		return 0, nil, errors.New("failed to decode response")
	}

	h, err := decodeHeader(&b)
	if err != nil {
		return 0, nil, err
	}

	b = b[:h.bodyLength]
	body, err := decodeBody(&b)
	return h.requestId, body, err
}

func decodeHeader(b *[]byte) (*CubeHeader, error) {
	if len(*b) < 3*intSize {
		return nil, errors.New("byte slice is too short for header")
	}
	header := CubeHeader{}

	var err error
	header.svcId, err = extractInt(b)
	if err != nil {
		return nil, err
	}

	header.bodyLength, err = extractInt(b)
	if err != nil {
		return nil, err
	}

	header.requestId, err = extractRId(b)
	if err != nil {
		return nil, err
	}

	return &header, nil
}

func decodeBody(b *[]byte) (*CubeResponseBody, error) {
	rc, err := extractInt(b)
	if err != nil {
		return nil, err
	}

	body := CubeResponseBody{ReturnCode: rc}
	if rc != ok {
		body.ErrString, err = extractString(b)
		if err != nil {
			return nil, err
		}
		return &body, nil
	}

	body.ClientId, err = extractString(b)
	if err != nil {
		return nil, err
	}

	body.ClientType, err = extractInt(b)
	if err != nil {
		return nil, err
	}

	body.Username, err = extractString(b)
	if err != nil {
		return nil, err
	}

	body.ExpiresIn, err = extractInt(b)
	if err != nil {
		return nil, err
	}

	body.UserId, err = extractInt64(b)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func extractInt(b *[]byte) (Int, error) {
	if b == nil {
		return 0, errors.New("empty ptr passed to extract int")
	}

	if len(*b) < intSize {
		return 0, errors.New("too short data passed for int extraction")
	}

	defer func() {
		temp := *b
		*b = temp[intSize:]
	}()

	return Int(binary.LittleEndian.Uint32(*b)), nil
}

func extractInt64(b *[]byte) (int64, error) {
	if b == nil {
		return 0, errors.New("empty ptr passed to extract int")
	}

	if len(*b) < 8 {
		return 0, errors.New("too short data passed for int64 extraction")
	}

	defer func() {
		temp := *b
		*b = temp[8:]
	}()

	return int64(binary.LittleEndian.Uint64(*b)), nil
}

func extractRId(b *[]byte) (RId, error) {
	v, err := extractInt(b)
	return RId(v), err
}

func extractString(b *[]byte) (string, error) {
	if b == nil {
		return "", errors.New("empty ptr passed to extract string")
	}

	if len(*b) < 4 {
		return "", errors.New("too short data passed for string len extraction")
	}

	l, err := extractInt(b)
	if err != nil {
		return "", err
	}

	// as int >= int32, it's correct to convert l to int from int32
	if int(l) > len(*b) {
		return "", errors.New("too short data passed for string extraction")
	}

	defer func() {
		temp := *b
		*b = temp[l:]
	}()

	return string((*b)[:l]), nil
}
