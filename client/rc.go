package client

const (
	ok Int = iota
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

func getErrorMessageByCode(code Int) string {
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
