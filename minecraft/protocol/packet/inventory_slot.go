package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// InventorySlot is sent by the server to update a single slot in one of the inventory windows that the client
// currently has opened. Usually this is the main inventory, but it may also be the off hand or, for example,
// a chest inventory.
//
// Added: v1.11.1
type InventorySlot struct {
	// WindowID is the ID of the window that the packet modifies. It must point to one of the windows that the
	// client currently has opened.
	//
	// Added: v1.11.1
	WindowID uint32
	// Slot is the index of the slot that the packet modifies. The new item will be set to the slot at this
	// index.
	//
	// Added: v1.11.1
	Slot uint32
	// Container is the protocol.FullContainerName that describes the container that the content is for.
	//
	// Added: v1.21.30
	// Changed: v1.26.20, encoded as an optional full container name instead of a plain full container name.
	// Changed: v1.26.20.26, encoded as an optional full container name in the reworked container attachment layout.
	Container protocol.Optional[protocol.FullContainerName]
	// StorageItem is the item that is acting as the storage container for the inventory. If the inventory is
	// not a dynamic container then this field should be left empty. When set, only the item type is used by
	// the client and none of the other stack info.
	//
	// Added: v1.21.40
	// Changed: v1.26.20, encoded as an optional storage item instead of a plain item instance.
	// Changed: v1.26.20.26, encoded as an optional storage item in the reworked container attachment layout.
	StorageItem protocol.Optional[protocol.ItemInstance]
	// NewItem is the item to be put in the slot at Slot. It will overwrite any item that may currently
	// be present in that slot.
	//
	// Added: v1.11.1
	NewItem protocol.ItemInstance
}

// ID ...
func (*InventorySlot) ID() uint32 {
	return IDInventorySlot
}

func (pk *InventorySlot) Marshal(io protocol.IO) {
	io.Varuint32(&pk.WindowID)
	io.Varuint32(&pk.Slot)
	if io.Protocol() >= protocol.Protocol1v26v20v26 {
		protocol.OptionalMarshaler(io, &pk.Container)
		protocol.OptionalFunc(io, &pk.StorageItem, io.ItemInstanceNew)
		io.ItemInstanceNew(&pk.NewItem)
	} else {
		var container protocol.FullContainerName
		protocol.Single(io, &container)
		pk.Container = protocol.Option(container)

		var storageItem protocol.ItemInstance
		io.ItemInstance(&storageItem)
		pk.StorageItem = protocol.Option(storageItem)

		io.ItemInstance(&pk.NewItem)
	}
}
