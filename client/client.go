package client

const TokenVerifySvcMsg = 0x00000001

type returnCode int

const (
	ok            returnCode = 0x00000000
	tokenNotFound returnCode = 0x00000001
	dbError       returnCode = 0x00000002
	unknownMsg    returnCode = 0x00000003
	badPacket     returnCode = 0x00000004
	badClient     returnCode = 0x00000005
	badScope      returnCode = 0x00000006
)
