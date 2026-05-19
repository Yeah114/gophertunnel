package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// ActorPickRequest is sent by the client when it tries to pick an entity, so that it gets a spawn egg which
// can spawn that entity.
//
// Added: v1.16
type ActorPickRequest struct {
	// EntityUniqueID is the unique ID of the entity that was attempted to be picked. The server must find the
	// type of that entity and provide the correct spawn egg to the player.
	//
	// Added: v1.16
	EntityUniqueID int64
	// HotBarSlot is the held hot bar slot of the player at the time of trying to pick the entity. If empty,
	// the resulting spawn egg should be put into this slot.
	//
	// Added: v1.16
	HotBarSlot byte
	// WithData is true if the pick request requests the entity metadata.
	//
	// Added: v1.16
	WithData bool
}

// ID ...
func (*ActorPickRequest) ID() uint32 {
	return IDActorPickRequest
}

func (pk *ActorPickRequest) Marshal(io protocol.IO) {
	io.Int64(&pk.EntityUniqueID)
	io.Uint8(&pk.HotBarSlot)
	io.Bool(&pk.WithData)
}
