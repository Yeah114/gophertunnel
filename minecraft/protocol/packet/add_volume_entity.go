package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddVolumeEntity sends a volume entity's definition and metadata from server to client.
//
// Added: v1.17.30
type AddVolumeEntity struct {
	// EntityRuntimeID is the runtime ID of the volume. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	//
	// Added: v1.17.30
	EntityRuntimeID uint32
	// EntityMetadata is a map of entity metadata, which includes flags and data properties that alter in
	// particular the way the volume functions or looks.
	//
	// Added: v1.17.30
	EntityMetadata map[string]any
	// EncodingIdentifier is the unique identifier for the volume. It must be of the form 'namespace:name', where
	// namespace cannot be 'minecraft'.
	//
	// Added: v1.18.10
	EncodingIdentifier string
	// InstanceIdentifier is the identifier of a fog definition.
	//
	// Added: v1.18.10
	InstanceIdentifier string
	// Bounds represent the volume's bounds. The first value is the minimum bounds, and the second value is the
	// maximum bounds.
	//
	// Added: v1.18.30
	Bounds [2]protocol.BlockPos
	// Dimension is the dimension in which the volume exists.
	//
	// Added: v1.18.30
	Dimension int32
	// EngineVersion is the engine version the entity is using, for example, '1.17.0'.
	//
	// Added: v1.17.30
	EngineVersion string
}

// ID ...
func (*AddVolumeEntity) ID() uint32 {
	return IDAddVolumeEntity
}

func (pk *AddVolumeEntity) Marshal(io protocol.IO) {
	io.Varuint32(&pk.EntityRuntimeID)
	io.NBT(&pk.EntityMetadata, nbt.NetworkLittleEndian)
	if io.Protocol() >= protocol.Protocol1v18v10 {
		io.String(&pk.EncodingIdentifier)
		io.String(&pk.InstanceIdentifier)
		if io.Protocol() >= protocol.Protocol1v18v30 {
			io.UBlockPos(&pk.Bounds[0])
			io.UBlockPos(&pk.Bounds[1])
			io.Varint32(&pk.Dimension)
		}
	}
	if io.Protocol() >= protocol.Protocol1v17v30 {
		io.String(&pk.EngineVersion)
	}
}
