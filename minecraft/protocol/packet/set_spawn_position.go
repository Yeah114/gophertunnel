package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	SpawnTypePlayer = iota
	SpawnTypeWorld
)

// SetSpawnPosition is sent by the server to update the spawn position of a player, for example when sleeping
// in a bed.
//
// Added: v1.12
type SetSpawnPosition struct {
	// SpawnType is the type of spawn to set. It is either SpawnTypePlayer or SpawnTypeWorld, and specifies
	// the behaviour of the spawn set. If SpawnTypeWorld is set, the position to which compasses will point is
	// also changed.
	//
	// Added: v1.12
	SpawnType int32
	// Position is the new position of the spawn that was set. If SpawnType is SpawnTypeWorld, compasses will
	// point to this position. As of 1.16, Position is always the position of the player.
	//
	// Added: v1.12
	// Changed: v1.16, no longer carries the world spawn position and instead always carries the player's spawn position.
	Position protocol.BlockPos
	// Dimension is the ID of the dimension that had its spawn updated. This is specifically relevant for
	// behaviour added in 1.16 such as the respawn anchor, which allows setting the spawn in a specific
	// dimension.
	//
	// Added: v1.16
	Dimension int32
	// SpawnPosition is a new field added in 1.16. It holds the spawn position of the world. This spawn
	// position is {-2147483648, -2147483648, -2147483648} for a default spawn position.
	//
	// Added: v1.16
	SpawnPosition protocol.BlockPos
}

// ID ...
func (*SetSpawnPosition) ID() uint32 {
	return IDSetSpawnPosition
}

func (pk *SetSpawnPosition) Marshal(io protocol.IO) {
	io.Varint32(&pk.SpawnType)
	io.UBlockPos(&pk.Position)
	io.Varint32(&pk.Dimension)
	io.UBlockPos(&pk.SpawnPosition)
}
