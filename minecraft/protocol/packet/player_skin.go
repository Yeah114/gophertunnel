package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
	"github.com/google/uuid"
)

// PlayerSkin is sent by the client to the server when it updates its own skin using the in-game skin picker.
// It is relayed by the server, or sent if the server changes the skin of a player on its own accord. Note
// that the packet can only be sent for players that are in the player list at the time of sending.
//
// Added: v1.12
type PlayerSkin struct {
	// UUID is the UUID of the player as sent in the Login packet when the client joined the server. It must
	// match this UUID exactly for the skin to show up on the player.
	//
	// Added: v1.12
	UUID uuid.UUID
	// Skin is the new skin to be applied on the player with the UUID in the field above. The skin, including
	// its animations, will be shown after sending it.
	//
	// Added: v1.12
	Skin protocol.Skin
	// NewSkinName no longer has a function: The field can be left empty at all times.
	//
	// Added: v1.12
	NewSkinName string
	// OldSkinName no longer has a function: The field can be left empty at all times.
	//
	// Added: v1.12
	OldSkinName string
}

// ID ...
func (*PlayerSkin) ID() uint32 {
	return IDPlayerSkin
}

func (pk *PlayerSkin) Marshal(io protocol.IO) {
	io.UUID(&pk.UUID)
	protocol.Single(io, &pk.Skin)
	io.String(&pk.NewSkinName)
	io.String(&pk.OldSkinName)
	io.Bool(&pk.Skin.Trusted)
}
