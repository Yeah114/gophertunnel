package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PartyChanged is sent by the client to the server to indicate that the player's party ID has changed.
type PartyChanged struct {
	// PartyID is the optional party identifier.
	PartyID protocol.Optional[string]
	// PartyLeader is if the client is the new party leader or not. This field was added in v1.26.20.26.
	PartyLeader bool
}

// ID ...
func (*PartyChanged) ID() uint32 {
	return IDPartyChanged
}

func (pk *PartyChanged) Marshal(io protocol.IO) {
	protocol.OptionalFunc(io, &pk.PartyID, io.String)
	if io.Protocol() >= protocol.Protocol1v26v20v26 {
		io.Bool(&pk.PartyLeader)
	}
}
