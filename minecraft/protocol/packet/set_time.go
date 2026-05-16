package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SetTime is sent by the server to update the current time client-side. The client actually advances time
// client-side by itself, so this packet does not need to be sent each tick. It is merely a means of
// synchronising time between server and client.
//
// Added: v1.11.1
type SetTime struct {
	// Time is the current time. The time is not limited to 24000 (time of day), but continues progressing
	// after that.
	//
	// Added: v1.11.1
	Time int32
}

// ID ...
func (*SetTime) ID() uint32 {
	return IDSetTime
}

func (pk *SetTime) Marshal(io protocol.IO) {
	io.Varint32(&pk.Time)
}
