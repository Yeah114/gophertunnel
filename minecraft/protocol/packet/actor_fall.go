package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

// ActorFall is sent by the client when an actor falls.
//
// Added: v1.11.1
type ActorFall struct {
	// EntityRuntimeID is the runtime ID of the falling actor.
	//
	// Added: v1.11.1
	EntityRuntimeID uint64
	// FallDistance is the distance the actor fell.
	//
	// Added: v1.11.1
	FallDistance float32
	// Unknown is an extra flag written after the fall distance.
	//
	// Added: v1.11.1
	Unknown bool
}

// ID ...
func (*ActorFall) ID() uint32 {
	return IDActorFall
}

func (pk *ActorFall) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Float32(&pk.FallDistance)
	io.Bool(&pk.Unknown)
}
