package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ClientStartItemCooldown is sent by the server to the client to initiate a cooldown on an item.
//
// Added: v1.18.10
type ClientStartItemCooldown struct {
	// Category is the category of the item to start the cooldown on.
	//
	// Added: v1.18.10
	Category string
	// Duration is the duration of ticks the cooldown should last.
	//
	// Added: v1.18.10
	Duration int32
}

// ID ...
func (*ClientStartItemCooldown) ID() uint32 {
	return IDClientStartItemCooldown
}

func (pk *ClientStartItemCooldown) Marshal(io protocol.IO) {
	io.String(&pk.Category)
	io.Varint32(&pk.Duration)
}
