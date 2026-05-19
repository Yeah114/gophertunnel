package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/nbt"
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// SyncActorProperty is an alternative to synced actor data.
//
// Added: v1.17.0
type SyncActorProperty struct {
	// PropertyData ...
	//
	// Added: v1.17.0
	PropertyData map[string]any
}

// ID ...
func (*SyncActorProperty) ID() uint32 {
	return IDSyncActorProperty
}

func (pk *SyncActorProperty) Marshal(io protocol.IO) {
	io.NBT(&pk.PropertyData, nbt.NetworkLittleEndian)
}
