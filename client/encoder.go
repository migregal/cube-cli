package client

import (
	"encoding/binary"
	"errors"
	"math/rand"
)

type Encoder struct {
	value []byte
}

func (e Encoder) FormatRequest(request *CubeRequest) (int, []byte, error) {
	if request == nil {
		return 0, nil, errors.New("empty ptr for request passed to encoder")
	}

	b, err := formatBody(request)
	if err != nil {
		return 0, nil, err
	}

	reqId := rand.Int()
	header, err := formatHeader(request.SvcId, len(b), reqId)
	if err != nil {
		return 0, nil, err
	}
	b = append(header, b...)

	return reqId, b, nil
}

func formatHeader(id, bodyLen, reqId int) ([]byte, error) {
	b := make([]byte, 0, 12)
	if err := appendInt(&b, id); err != nil {
		return nil, err
	}
	if err := appendInt(&b, bodyLen); err != nil {
		return nil, err
	}
	if err := appendInt(&b, reqId); err != nil {
		return nil, err
	}
	return b, nil
}

func formatBody(request *CubeRequest) (b []byte, err error) {
	if request == nil {
		return nil, errors.New("empty ptr for request passed to body encoder")
	}

	b = make([]byte, 0, 4+len(request.Token)+len(request.Scope))

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

func appendString(b *[]byte, value string) error {
	if b == nil {
		return errors.New("empty ptr passed to append string")
	}
	if err := appendInt(b, len(value)); err != nil {
		return err
	}
	*b = append(*b, []byte(value)...)
	return nil
}

func appendInt(b *[]byte, value int) error {
	if b == nil {
		return errors.New("empty ptr passed to append int")
	}
	temp := make([]byte, 4)
	binary.LittleEndian.PutUint32(temp, uint32(value))
	*b = append(*b, temp...)
	return nil
}
