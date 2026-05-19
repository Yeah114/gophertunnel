package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

const (
	MovementEffectTypeGlideBoost = iota
)

// MovementEffect is sent by the server to the client to update specific movement effects to allow the client
// to predict its movement. For example, fireworks used during gliding will send this packet to tell the
// client the exact duration of the boost.
//
// Added: v1.21.40
type MovementEffect struct {
	// EntityRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	//
	// Added: v1.21.40
	EntityRuntimeID uint64
	// Type is the type of movement effect being updated. It is one of the constants found above.
	//
	// Added: v1.21.40
	Type int32
	// Duration is the duration of the effect, measured in ticks.
	//
	// Added: v1.21.40
	Duration int32
	// Tick is the server tick at which the packet was sent. It is used in relation to CorrectPlayerMovePrediction.
	//
	// Added: v1.21.40
	Tick uint64
}

// ID ...
func (*MovementEffect) ID() uint32 {
	return IDMovementEffect
}

func (pk *MovementEffect) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Varint32(&pk.Type)
	io.Varint32(&pk.Duration)
	io.Varuint64(&pk.Tick)
}
