package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// ClientMovementPredictionSync is sent by the client to the server periodically if the client has received
// movement corrections from the server, containing information about client-predictions that are relevant
// to movement.
//
// Added: v1.21.60
type ClientMovementPredictionSync struct {
	// ActorFlags is a bitset of all the flags that are currently set for the client.
	//
	// Added: v1.21.60
	ActorFlags protocol.Bitset
	// BoundingBoxScale is the scale of the client's bounding box.
	//
	// Added: v1.21.60
	BoundingBoxScale float32
	// BoundingBoxWidth is the width of the client's bounding box.
	//
	// Added: v1.21.60
	BoundingBoxWidth float32
	// BoundingBoxHeight is the height of the client's bounding box.
	//
	// Added: v1.21.60
	BoundingBoxHeight float32
	// MovementSpeed is the movement speed attribute or 0 if not set.
	//
	// Added: v1.21.60
	MovementSpeed float32
	// UnderwaterMovementSpeed is the underwater movement speed attribute or 0 if not set.
	//
	// Added: v1.21.60
	UnderwaterMovementSpeed float32
	// LavaMovementSpeed is the lava movement speed attribute or 0 if not set.
	//
	// Added: v1.21.60
	LavaMovementSpeed float32
	// JumpStrength is the jump strength attribute or 0 if not set.
	//
	// Added: v1.21.60
	JumpStrength float32
	// Health is the health attribute or 0 if not set.
	//
	// Added: v1.21.60
	Health float32
	// Hunger is the hunger attribute or 0 if not set.
	//
	// Added: v1.21.60
	Hunger float32
	// FrictionModifier is the friction modifier movement attribute.
	//
	// Added: v1.26.20
	FrictionModifier float32
	// Bounciness is the bounciness movement attribute.
	//
	// Added: v1.26.20
	Bounciness float32
	// AirDragModifier is the air drag modifier movement attribute.
	//
	// Added: v1.26.20
	AirDragModifier float32
	// EntityRuntimeID is the runtime ID of the entity that the prediction data applies to.
	//
	// Added: v1.21.60
	EntityRuntimeID uint64
	// Flying specifies if the client is currently flying.
	//
	// Added: v1.21.80
	Flying bool
}

// ID ...
func (*ClientMovementPredictionSync) ID() uint32 {
	return IDClientMovementPredictionSync
}

func (pk *ClientMovementPredictionSync) Marshal(io protocol.IO) {
	io.Bitset(&pk.ActorFlags, protocol.EntityDataFlagCount)
	io.Float32(&pk.BoundingBoxScale)
	io.Float32(&pk.BoundingBoxWidth)
	io.Float32(&pk.BoundingBoxHeight)
	io.Float32(&pk.MovementSpeed)
	io.Float32(&pk.UnderwaterMovementSpeed)
	io.Float32(&pk.LavaMovementSpeed)
	io.Float32(&pk.JumpStrength)
	io.Float32(&pk.Health)
	io.Float32(&pk.Hunger)
	if io.Protocol() >= protocol.Protocol1v26v20 {
		io.Float32(&pk.FrictionModifier)
		io.Float32(&pk.Bounciness)
		io.Float32(&pk.AirDragModifier)
	}
	io.Varuint64(&pk.EntityRuntimeID)
	if io.Protocol() >= protocol.Protocol1v21v80 {
		io.Bool(&pk.Flying)
	}
}
