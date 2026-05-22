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
	// Changed: v1.21.110, encoded as an array of armour slot and damage pairs instead of a bitset and damage list.
	List []protocol.PlayerArmourDamageEntry
}

// ID ...
func (pk *PlayerArmourDamage) ID() uint32 {
	return IDPlayerArmourDamage
}

func (pk *PlayerArmourDamage) Marshal(io protocol.IO) {
	if io.Protocol() >= protocol.Protocol1v21v110 {
		protocol.Slice(io, &pk.List)
		return
	}

	var flags uint8
	maxSlot := int32(PlayerArmourDamageFlagBoots)
	if io.Protocol() >= protocol.Protocol1v21v20 {
		maxSlot = PlayerArmourDamageFlagBody
	}
	for _, entry := range pk.List {
		if entry.ArmourSlot < 0 || entry.ArmourSlot > maxSlot {
			continue
		}
		flags |= 1 << uint8(entry.ArmourSlot)
	}
	io.Uint8(&flags)
	for _, entry := range pk.List {
		if entry.ArmourSlot < 0 || entry.ArmourSlot > maxSlot || flags&(1<<uint8(entry.ArmourSlot)) == 0 {
			continue
		}
		damage := int32(entry.Damage)
		io.Varint32(&damage)
		entry.Damage = int16(damage)
	}
}
