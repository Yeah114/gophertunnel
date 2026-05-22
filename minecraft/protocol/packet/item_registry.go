package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/nbt"
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// ItemRegistry is sent by the server to send the client a list of available items and attach client-side
// components to a custom item. This packet was formerly known as the ItemComponent packet before 1.21.60,
// which did not include item definitions but only the components.
//
// Added: v1.16.100
// Changed: v1.21.60, renamed from ItemComponent and expanded to include item definitions.
type ItemRegistry struct {
	// Items is a list of all items with their legacy IDs which are available in the game. Failing to send any
	// of the items that are in the game will crash mobile clients. Any custom components are also attached to
	// the items in this list.
	//
	// Added: v1.16.100
	// Changed: v1.21.60, expanded from component data to full item entries with runtime IDs and versions.
	Items []protocol.ItemEntry
}

// ID ...
func (*ItemRegistry) ID() uint32 {
	return IDItemRegistry
}

func (pk *ItemRegistry) Marshal(io protocol.IO) {
	if io.Protocol() < protocol.Protocol1v21v60 {
		protocol.FuncSlice(io, &pk.Items, func(x *protocol.ItemEntry) {
			io.String(&x.Name)
			io.NBT(&x.Data, nbt.NetworkLittleEndian)
		})
		return
	}
	protocol.Slice(io, &pk.Items)
}
