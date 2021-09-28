package client

import "unsafe"

type CubeHeader struct {
	svcId      Int
	bodyLength Int
	requestId  RId
}

const headerSize = int(unsafe.Sizeof(CubeHeader{}))
