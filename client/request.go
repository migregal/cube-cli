package client

type CubeRequest struct {
	Token string
	Scope string
	SvcId int32
}

type CubeHeader struct {
	svcId      int32
	bodyLength int32
	requestId  int32
}
