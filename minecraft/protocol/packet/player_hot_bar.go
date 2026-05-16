package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PlayerHotBar is sent by the server to the client. It used to be used to link hot bar slots of the player to
// actual slots in the inventory, but as of 1.2, this was changed and hot bar slots are no longer a free
// floating part of the inventory.
// Since 1.2, the packet has been re-purposed, but its new functionality is not clear.
//
// Added: v1.11.1
type PlayerHotBar struct {
	// SelectedHotBarSlot is the selected hotbar slot index that the client should treat as active.
	//
	// Added: v1.11.1
	SelectedHotBarSlot uint32
	// WindowID is the inventory window that the selected hotbar slot belongs to.
	//
	// Added: v1.11.1
	WindowID byte
	// SelectHotBarSlot specifies if the client should actually switch to the selected hotbar slot.
	//
	// Added: v1.11.1
	SelectHotBarSlot bool
}

// ID ...
func (*PlayerHotBar) ID() uint32 {
	return IDPlayerHotBar
}

func (pk *PlayerHotBar) Marshal(io protocol.IO) {
	io.Varuint32(&pk.SelectedHotBarSlot)
	io.Uint8(&pk.WindowID)
	if io.Protocol() == protocol.Protocol1v2v10 {
		legacyUnknown := uint32(0)
		io.Varuint32(&legacyUnknown)
	}
	io.Bool(&pk.SelectHotBarSlot)
}
