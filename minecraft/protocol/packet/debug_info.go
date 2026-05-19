package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// DebugInfo is a packet sent by the server to the client. It does not seem to do anything when sent to the
// normal client in 1.16.
//
// Added: v1.16
type DebugInfo struct {
	// PlayerUniqueID is the unique ID of the player that the packet is sent to.
	//
	// Added: v1.16
	PlayerUniqueID int64
	// Data is the debug data.
	//
	// Added: v1.16
	Data []byte
}

// ID ...
func (*DebugInfo) ID() uint32 {
	return IDDebugInfo
}

func (pk *DebugInfo) Marshal(io protocol.IO) {
	io.Varint64(&pk.PlayerUniqueID)
	io.ByteSlice(&pk.Data)
}
