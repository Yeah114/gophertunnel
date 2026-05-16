package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// RemoveObjective is sent by the server to remove a scoreboard objective. It is used to stop showing a
// scoreboard to a player.
//
// Added: v1.12
type RemoveObjective struct {
	// ObjectiveName is the name of the objective that the scoreboard currently active has. This name must
	// be identical to the one sent in the SetDisplayObjective packet.
	//
	// Added: v1.12
	ObjectiveName string
}

// ID ...
func (*RemoveObjective) ID() uint32 {
	return IDRemoveObjective
}

func (pk *RemoveObjective) Marshal(io protocol.IO) {
	io.String(&pk.ObjectiveName)
}
