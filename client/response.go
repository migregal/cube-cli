package client

type CubeResponseBody struct {
	ReturnCode int32

	ClientId   string
	ClientType int32
	Username   string
	ExpiresIn  int32
	UserId     int64

	ErrString string
}
