package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SetDifficulty is sent by the server to update the client-side difficulty of the client. The actual effect
// of this packet on the client isn't very significant, as the difficulty is handled server-side.
//
// Added: v1.11.1
type SetDifficulty struct {
	// Difficulty is the new difficulty that the world has.
	//
	// Added: v1.11.1
	Difficulty uint32
}

// ID ...
func (*SetDifficulty) ID() uint32 {
	return IDSetDifficulty
}

func (pk *SetDifficulty) Marshal(io protocol.IO) {
	io.Varuint32(&pk.Difficulty)
}
