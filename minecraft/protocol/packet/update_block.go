package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	BlockUpdateNeighbours = 1 << iota
	BlockUpdateNetwork
	BlockUpdateNoGraphics
	BlockUpdatePriority
)

// UpdateBlock is sent by the server to update a block client-side, without resending the entire chunk that
// the block is located in. It is particularly useful for small modifications like block breaking/placing.
//
// Added: v1.12
type UpdateBlock struct {
	// Position is the block position at which a block is updated.
	//
	// Added: v1.12
	Position protocol.BlockPos
	// NewBlockRuntimeID is the runtime ID of the block that is placed at Position after sending the packet
	// to the client.
	//
	// Added: v1.12
	NewBlockRuntimeID uint32
	// LegacyBlockID is the numeric block ID written before block runtime IDs were introduced.
	//
	// Removed: v1.2.13
	LegacyBlockID uint32
	// LegacyBlockData is the block data written alongside LegacyBlockID.
	//
	// Removed: v1.2.13
	LegacyBlockData uint32
	// Flags is a combination of flags that specify the way the block is updated client-side. It is a
	// combination of the flags above, but typically sending only the BlockUpdateNetwork flag is sufficient.
	//
	// Added: v1.12
	Flags uint32
	// Layer is the world layer on which the block is updated. For most blocks, this is the first layer, as
	// that layer is the default layer to place blocks on, but for blocks inside of each other, this differs.
	//
	// Added: v1.12
	Layer uint32
}

// ID ...
func (*UpdateBlock) ID() uint32 {
	return IDUpdateBlock
}

func (pk *UpdateBlock) Marshal(io protocol.IO) {
	io.UBlockPos(&pk.Position)
	if io.Protocol() > protocol.Protocol1v2v10 {
		io.Varuint32(&pk.NewBlockRuntimeID)
		io.Varuint32(&pk.Flags)
	} else {
		io.Varuint32(&pk.LegacyBlockID)
		legacyBlockData := 0xb0 | pk.LegacyBlockData&0xf
		io.Varuint32(&legacyBlockData)
	}
	if io.Protocol() > protocol.Protocol1v2v13v11 {
		io.Varuint32(&pk.Layer)
	}
}
