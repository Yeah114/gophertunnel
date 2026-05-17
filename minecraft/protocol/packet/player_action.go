package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PlayerAction is sent by the client when it executes any action, for example starting to sprint, swim,
// starting the breaking of a block, dropping an item, etc.
//
// Added: v1.11.1
type PlayerAction struct {
	// EntityRuntimeID is the runtime ID of the player. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	//
	// Added: v1.11.1
	EntityRuntimeID uint64
	// ActionType is the ID of the action that was executed by the player. It is one of the constants that may
	// be found in protocol/player.go.
	//
	// Added: v1.11.1
	ActionType int32
	// BlockPosition is the position of the target block, if the action with the ActionType set concerned a
	// block. If that is not the case, the block position will be zero.
	//
	// Added: v1.11.1
	BlockPosition protocol.BlockPos
	// ResultPosition is the position of the action's result. When a UseItemOn action is sent, this is the position of
	// the block clicked, but when a block is placed, this is the position at which the block will be placed.
	//
	// Added: v1.19.0.29
	ResultPosition protocol.BlockPos
	// BlockFace is the face of the target block that was touched. If the action with the ActionType set
	// concerned a block. If not, the face is always 0.
	//
	// Added: v1.11.1
	BlockFace int32
}

// ID ...
func (*PlayerAction) ID() uint32 {
	return IDPlayerAction
}

func (pk *PlayerAction) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Varint32(&pk.ActionType)
	io.UBlockPos(&pk.BlockPosition)
	if io.Protocol() >= protocol.Protocol1v19v0v29 {
		io.UBlockPos(&pk.ResultPosition)
	}
	io.Varint32(&pk.BlockFace)
}
