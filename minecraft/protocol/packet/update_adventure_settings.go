package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// UpdateAdventureSettings is a packet sent from the server to the client to update the adventure settings of the player.
// It, along with the UpdateAbilities packet, are replacements of the AdventureSettings packet since v1.19.10.
//
// Added: v1.19.10
type UpdateAdventureSettings struct {
	// NoPvM is a boolean indicating whether the player is allowed to fight mobs or not.
	//
	// Added: v1.19.10
	NoPvM bool
	// NoMvP is a boolean indicating whether mobs are allowed to fight the player or not. It is unclear why this is sent
	// to the client.
	//
	// Added: v1.19.10
	NoMvP bool
	// ImmutableWorld is a boolean indicating whether the player is allowed to modify the world or not.
	//
	// Added: v1.19.10
	ImmutableWorld bool
	// ShowNameTags is a boolean indicating whether player name tags are shown or not.
	//
	// Added: v1.19.10
	ShowNameTags bool
	// AutoJump is a boolean indicating whether the player is allowed to jump automatically or not.
	//
	// Added: v1.19.10
	AutoJump bool
}

// ID ...
func (*UpdateAdventureSettings) ID() uint32 {
	return IDUpdateAdventureSettings
}

func (pk *UpdateAdventureSettings) Marshal(io protocol.IO) {
	io.Bool(&pk.NoPvM)
	io.Bool(&pk.NoMvP)
	io.Bool(&pk.ImmutableWorld)
	io.Bool(&pk.ShowNameTags)
	io.Bool(&pk.AutoJump)
}
