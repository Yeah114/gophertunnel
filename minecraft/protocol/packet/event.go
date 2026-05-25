package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// Event is sent by the server to send an event with additional data. It is typically sent to the client for
// telemetry reasons, much like the SimpleEvent packet.
//
// Added: v1.11.1
type Event struct {
	// EntityRuntimeID is the runtime ID of the player. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	//
	// Added: v1.11.1
	// Changed: v1.21.20, encoded as a signed varint64 instead of an unsigned varuint64.
	EntityRuntimeID int64
	// UsePlayerID ...
	// TODO: Figure out what UsePlayerID is for.
	//
	// Added: v1.17.0
	// Changed: v1.21.130.28, encoded as a bool instead of a uint8.
	UsePlayerID bool
	// Event is the event that is transmitted.
	//
	// Added: v1.17.0
	// Changed: v1.21.130.28, followed by an event ordinal before the event payload.
	Event protocol.Event
}

// ID ...
func (*Event) ID() uint32 {
	return IDEvent
}

func (pk *Event) Marshal(io protocol.IO) {
	io.Varint64(&pk.EntityRuntimeID)
	io.EventType(&pk.Event)
	if io.Protocol() >= protocol.Protocol1v21v130v28 {
		io.Bool(&pk.UsePlayerID)
		io.EventOrdinal(&pk.Event)
	} else {
		usePlayerID := uint8(0)
		if pk.UsePlayerID {
			usePlayerID = 1
		}
		io.Uint8(&usePlayerID)
		pk.UsePlayerID = usePlayerID != 0
	}
	pk.Event.Marshal(io)
}
