package client

import (
	"unsafe"
)

type Int int32

const intSize = int(unsafe.Sizeof(Int(0)))

type Int64 int64

const intSize64 = int(unsafe.Sizeof(Int64(0)))

type RId Int

const ridSize = int(unsafe.Sizeof(RId(0)))
