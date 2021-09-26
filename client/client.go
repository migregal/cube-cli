package client

const TokenVerifySvcMsg = 0x00000001

const (
	ok            int32 = 0x00000000
	tokenNotFound int32 = 0x00000001
	dbError       int32 = 0x00000002
	unknownMsg    int32 = 0x00000003
	badPacket     int32 = 0x00000004
	badClient     int32 = 0x00000005
	badScope      int32 = 0x00000006
)
