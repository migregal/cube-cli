package client

import (
	"bytes"
	"fmt"
)

type CubeResponseBody struct {
	ReturnCode Int

	ClientId   string
	ClientType Int
	Username   string
	ExpiresIn  Int
	UserId     Int64

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
