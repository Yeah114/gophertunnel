package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// RemoveVolumeEntity indicates a volume entity to be removed from server to client.
//
// Added: v1.17.0
type RemoveVolumeEntity struct {
	// EntityRuntimeID is the runtime ID of the volume entity that should be removed.
	//
	// Added: v1.17.0
	EntityRuntimeID uint32
	// Dimension is the dimension that the volume entity belongs to.
	//
	// Added: v1.17.0
	Dimension int32
}

// ID ...
func (*RemoveVolumeEntity) ID() uint32 {
	return IDRemoveVolumeEntity
}

func (pk *RemoveVolumeEntity) Marshal(io protocol.IO) {
	io.Varuint32(&pk.EntityRuntimeID)
	io.Varint32(&pk.Dimension)
}
