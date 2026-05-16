package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// InventoryContent is sent by the server to update the full content of a particular inventory. It is usually
// sent for the main inventory of the player, but also works for other inventories that are currently opened
// by the player.
//
// Added: v1.11.1
type InventoryContent struct {
	// WindowID is the ID that identifies one of the windows that the client currently has opened, or one of
	// the consistent windows such as the main inventory.
	//
	// Added: v1.11.1
	WindowID uint32
	// Content is the new content of the inventory. The length of this slice must be equal to the full size of
	// the inventory window updated.
	//
	// Added: v1.11.1
	Content []protocol.ItemInstance
	// Container is the protocol.FullContainerName that describes the container that the content is for.
	//
	// Added: v1.21.30
	// Changed: v1.21.40, replaced the DynamicWindowID field layout introduced in v1.21.20.
	Container protocol.FullContainerName
	// StorageItem is the item that is acting as the storage container for the inventory. If the inventory is
	// not a dynamic container then this field should be left empty. When set, only the item type is used by
	// the client and none of the other stack info.
	//
	// Added: v1.21.40
	// Changed: v1.21.40, replaced the DynamicContainerSize field introduced in v1.21.30.
	StorageItem protocol.ItemInstance
}

// ID ...
func (*InventoryContent) ID() uint32 {
	return IDInventoryContent
}

func (pk *InventoryContent) Marshal(io protocol.IO) {
	io.Varuint32(&pk.WindowID)
	protocol.FuncSlice(io, &pk.Content, io.ItemInstance)
	protocol.Single(io, &pk.Container)
	io.ItemInstance(&pk.StorageItem)
}
