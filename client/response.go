package client

import (
	"bytes"
	"fmt"
)

const (
	ok int32 = iota
	tokenNotFound
	dbError
	unknownMsg
	badPacket
	badClient
	badScope
)

const (
	msgOk            = "ok"
	msgTokenNotFound = "token not found"
	msgDbError       = "db error"
	msgUnknownMsg    = "unknown svc msg"
	msgBadPacket     = "bad packet"
	msgBadClient     = "bad client"
	msgBadScope      = "bad scope"
	msgUnknownError  = "unknown error"
)

type CubeResponseBody struct {
	ReturnCode int32

	ClientId   string
	ClientType int32
	Username   string
	ExpiresIn  int32
	UserId     int64

	ErrString string
}

func (crb *CubeResponseBody) ToString() string {
	if crb.ReturnCode != ok {
		return crb.buildErrorString()
	}

	return crb.buildOkString()
}

func (crb *CubeResponseBody) buildErrorString() string {
	var buffer bytes.Buffer

	buffer.WriteString(
		fmt.Sprintf(
			"error: %s\n",
			getErrorMessageByCode(crb.ReturnCode),
		),
	)
	buffer.WriteString(fmt.Sprintf("message: %s\n", crb.ErrString))

	return buffer.String()
}

func (crb *CubeResponseBody) buildOkString() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("client_id: %s\n", crb.ClientId))
	buffer.WriteString(fmt.Sprintf("client_type: %d\n", crb.ClientType))
	buffer.WriteString(fmt.Sprintf("expires_in: %d\n", crb.ExpiresIn))
	buffer.WriteString(fmt.Sprintf("user_id: %d\n", crb.UserId))
	buffer.WriteString(fmt.Sprintf("username: %s\n", crb.Username))

	return buffer.String()
}

func getErrorMessageByCode(code int32) string {
	switch code {
	case ok:
		return msgOk
	case tokenNotFound:
		return msgTokenNotFound
	case dbError:
		return msgDbError
	case unknownMsg:
		return msgUnknownMsg
	case badPacket:
		return msgBadPacket
	case badClient:
		return msgBadClient
	case badScope:
		return msgBadScope
	default:
		return msgUnknownError
	}
}
