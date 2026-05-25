package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

// ItemFrameDropItem is sent by the client when it drops an item from an item frame.
//
// Added: v1.11.1
type ItemFrameDropItem struct {
	// Position is the position of the item frame.
	//
	// Added: v1.11.1
	Position protocol.BlockPos
}

// ID ...
func (*ItemFrameDropItem) ID() uint32 {
	return IDItemFrameDropItem
}

func (pk *ItemFrameDropItem) Marshal(io protocol.IO) {
	io.BlockPos(&pk.Position)
}
