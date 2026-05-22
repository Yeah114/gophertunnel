package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// MobArmourEquipment is sent by the server to the client to update the armour an entity is wearing. It is
// sent for both players and other entities, such as zombies.
//
// Added: v1.12
type MobArmourEquipment struct {
	// EntityRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	//
	// Added: v1.12
	EntityRuntimeID uint64
	// Helmet is the equipped helmet of the entity. Items that are not wearable on the head will not be
	// rendered by the client. Unlike in Java Edition, blocks cannot be worn.
	//
	// Added: v1.12
	Helmet protocol.ItemInstance
	// Chestplate is the chestplate of the entity. Items that are not wearable as chestplate will not be
	// rendered.
	//
	// Added: v1.12
	Chestplate protocol.ItemInstance
	// Leggings is the item worn as leggings by the entity. Items not wearable as leggings will not be
	// rendered client-side.
	//
	// Added: v1.12
	Leggings protocol.ItemInstance
	// Boots is the item worn as boots by the entity. Items not wearable as boots will not be rendered.
	//
	// Added: v1.12
	Boots protocol.ItemInstance
	// Body is the item worn on the body of the entity. Items not wearable on the body will not be rendered.
	//
	// Added: v1.21.20
	Body protocol.ItemInstance
}

// ID ...
func (*MobArmourEquipment) ID() uint32 {
	return IDMobArmourEquipment
}

func (pk *MobArmourEquipment) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.ItemInstance(&pk.Helmet)
	io.ItemInstance(&pk.Chestplate)
	io.ItemInstance(&pk.Leggings)
	io.ItemInstance(&pk.Boots)
	if io.Protocol() >= protocol.Protocol1v21v20 {
		io.ItemInstance(&pk.Body)
	}
}
