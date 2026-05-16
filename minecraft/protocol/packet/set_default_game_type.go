package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SetDefaultGameType is sent by the client when it toggles the default game type in the settings UI, and is
// sent by the server when it actually changes the default game type, resulting in the toggle being changed
// in the settings UI.
//
// Added: v1.12
type SetDefaultGameType struct {
	// GameType is the new game type that is set. When sent by the client, this is the requested new default
	// game type.
	//
	// Added: v1.12
	GameType int32
}

// ID ...
func (*SetDefaultGameType) ID() uint32 {
	return IDSetDefaultGameType
}

func (pk *SetDefaultGameType) Marshal(io protocol.IO) {
	io.Varint32(&pk.GameType)
}
