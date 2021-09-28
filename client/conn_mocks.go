package client

import (
	"errors"
	"io"
	"time"
)

type EmptyWriteMockConn struct{}

func (c EmptyWriteMockConn) Write([]byte) (int, error) {
	return 0, nil
}

func (c EmptyWriteMockConn) Read([]byte) (int, error) {
	return 0, io.EOF
}

func (c EmptyWriteMockConn) Close() error {
	return nil
}

func (c EmptyWriteMockConn) SetDeadline(time.Time) error {
	return nil
}

type EmptyReadMockConn struct{}

func (c EmptyReadMockConn) Write(b []byte) (int, error) {
	return len(b), nil
}

func (c EmptyReadMockConn) Read(b []byte) (int, error) {
	return 0, errors.New("unable to read")
}

func (c EmptyReadMockConn) Close() error {
	return nil
}

func (c EmptyReadMockConn) SetDeadline(time.Time) error {
	return nil
}

type BrokenRespFmtMockConn struct {
	respLen int
}

func (c BrokenRespFmtMockConn) Write(b []byte) (int, error) {
	return len(b), nil
}

func (c BrokenRespFmtMockConn) Read([]byte) (int, error) {
	return c.respLen, io.EOF
}

func (c BrokenRespFmtMockConn) Close() error {
	return nil
}

func (c BrokenRespFmtMockConn) SetDeadline(time.Time) error {
	return nil
}

type FixedResponseMockConn struct {
	expected []byte
}

func (c FixedResponseMockConn) Write(b []byte) (int, error) {
	return len(b), nil
}

func (c FixedResponseMockConn) Read(b []byte) (int, error) {
	copy(b, c.expected)
	return len(c.expected), io.EOF
}

func (c FixedResponseMockConn) Close() error {
	return nil
}

func (c FixedResponseMockConn) SetDeadline(time.Time) error {
	return nil
}
