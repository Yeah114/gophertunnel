package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// ServerPresenceInfo is sent by the server to provide the client with presence info.
//
// Added: v1.26.20
type ServerPresenceInfo struct {
	// PresenceInfo is the presence info to set, or nothing to fall back to the default.
	//
	// Added: v1.26.20
	PresenceInfo protocol.Optional[protocol.PresenceInfo]
}

// ID ...
func (*ServerPresenceInfo) ID() uint32 {
	return IDServerPresenceInfo
}

func (pk *ServerPresenceInfo) Marshal(io protocol.IO) {
	protocol.OptionalMarshaler(io, &pk.PresenceInfo)
}
