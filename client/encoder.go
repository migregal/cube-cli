package client

import (
	"encoding/binary"
	"errors"
	"math/rand"
)

type Encoder struct{}

func (e Encoder) FormatRequest(request *CubeRequestBody) (RId, []byte, error) {
	if request == nil {
		return 0, nil, errors.New("empty ptr for request passed to encoder")
	}

	b, err := formatBody(request)
	if err != nil {
		return 0, nil, err
	}

	header := CubeHeader{
		svcId:      request.SvcId,
		bodyLength: Int(len(b)),
		requestId:  RId(rand.Int31()),
	}
	hb, err := formatHeader(&header)
	if err != nil {
		return 0, nil, err
	}
	b = append(hb, b...)

	return header.requestId, b, nil
}

func formatHeader(header *CubeHeader) ([]byte, error) {
	b := make([]byte, 0, headerSize)

	if err := appendInt(&b, header.svcId); err != nil {
		return nil, err
	}
	if err := appendInt(&b, header.bodyLength); err != nil {
		return nil, err
	}
	if err := appendRId(&b, header.requestId); err != nil {
		return nil, err
	}
	return b, nil
}

func formatBody(request *CubeRequestBody) (b []byte, err error) {
	if request == nil {
		return nil, errors.New("empty ptr for request passed to body encoder")
	}

	b = make([]byte, 0, intSize+len(request.Token)+len(request.Scope))

	if err = appendInt(&b, TokenVerifySvcMsg); err != nil {
		return nil, err
	}
	if err = appendString(&b, request.Token); err != nil {
		return nil, err
	}
	if err = appendString(&b, request.Scope); err != nil {
		return nil, err
	}

	return b, nil
}

func appendInt(b *[]byte, value Int) error {
	if b == nil {
		return errors.New("empty ptr passed to append int")
	}
	temp := make([]byte, intSize)
	binary.LittleEndian.PutUint32(temp, uint32(value))
	*b = append(*b, temp...)
	return nil
}

func appendRId(b *[]byte, value RId) error {
	return appendInt(b, Int(value))
}

func appendString(b *[]byte, value string) error {
	if b == nil {
		return errors.New("empty ptr passed to append string")
	}

	if err := appendInt(b, Int(len(value))); err != nil {
		return err
	}
	*b = append(*b, []byte(value)...)
	return nil
}
