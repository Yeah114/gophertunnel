package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PartyChanged is sent by the client to the server to indicate that the player's party ID has changed.
//
// Added: v1.26.10
type PartyChanged struct {
	// PartyInfo contains the optional party information payload sent by the client.
	//
	// Added: v1.26.10
	PartyInfo protocol.Optional[PartyInfo]
}

// ID ...
func (*PartyChanged) ID() uint32 {
	return IDPartyChanged
}

func (pk *PartyChanged) Marshal(io protocol.IO) {
	protocol.OptionalMarshaler(io, &pk.PartyInfo)
}

// PartyInfo represents the information of the client's role in a party.
//
// Added: v1.26.10
type PartyInfo struct {
	// PartyID is the party identifier.
	//
	// Added: v1.26.10
	PartyID string
	// PartyLeader is if the client is the new party leader or not.
	//
	// Added: v1.26.20.26
	PartyLeader bool
}

func (x *PartyInfo) Marshal(io protocol.IO) {
	io.String(&x.PartyID)
	if io.Protocol() >= protocol.Protocol1v26v20v26 {
		io.Bool(&x.PartyLeader)
	}
}
