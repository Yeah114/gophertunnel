package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

const (
	PlayerArmourDamageFlagHelmet = iota
	PlayerArmourDamageFlagChestplate
	PlayerArmourDamageFlagLeggings
	PlayerArmourDamageFlagBoots
	PlayerArmourDamageFlagBody
)

// PlayerArmourDamage is sent by the server to damage the armour of a player. It is a very efficient packet,
// but generally it's much easier to just send a slot update for the damaged armour.
//
// Added: v1.16
type PlayerArmourDamage struct {
	// List is a list of armour entries indicating which pieces of armour should receive damage.
	//
	// Added: v1.16
	// Changed: v1.21.20, may include body armour entries.
	List []protocol.PlayerArmourDamageEntry
}

// ID ...
func (pk *PlayerArmourDamage) ID() uint32 {
	return IDPlayerArmourDamage
}

func (pk *PlayerArmourDamage) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.List)
}
