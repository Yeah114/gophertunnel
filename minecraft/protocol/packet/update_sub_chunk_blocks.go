package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// UpdateSubChunkBlocks is essentially just UpdateBlock packet, however for a set of blocks in a sub-chunk.
//
// Added: v1.17.30
type UpdateSubChunkBlocks struct {
	// Position is the block position of the sub-chunk being referred to.
	//
	// Added: v1.17.30
	Position protocol.BlockPos
	// Blocks contains each updated block change entry.
	//
	// Added: v1.17.30
	Blocks []protocol.BlockChangeEntry
	// Extra contains each updated block change entry for the second layer, usually for waterlogged blocks.
	//
	// Added: v1.17.30
	Extra []protocol.BlockChangeEntry
}

// ID ...
func (*UpdateSubChunkBlocks) ID() uint32 {
	return IDUpdateSubChunkBlocks
}

func (pk *UpdateSubChunkBlocks) Marshal(io protocol.IO) {
	io.UBlockPos(&pk.Position)
	protocol.Slice(io, &pk.Blocks)
	protocol.Slice(io, &pk.Extra)
}
