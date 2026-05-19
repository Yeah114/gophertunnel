package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// RequestNetworkSettings is sent by the client to request network settings, such as compression, from the server.
//
// Added: v1.19.30
type RequestNetworkSettings struct {
	// ClientProtocol is the protocol version of the player. The player is disconnected if the protocol is
	// incompatible with the protocol of the server.
	//
	// Added: v1.19.30
	ClientProtocol int32
}

// ID ...
func (pk *RequestNetworkSettings) ID() uint32 {
	return IDRequestNetworkSettings
}

func (pk *RequestNetworkSettings) Marshal(io protocol.IO) {
	io.BEInt32(&pk.ClientProtocol)
}
