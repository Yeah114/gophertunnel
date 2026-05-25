package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
	"github.com/google/uuid"
)

const (
	CraftingEventTypeInventory = iota
	CraftingEventTypeCrafting
	CraftingEventTypeWorkbench
)

// CraftingEvent is sent by the client after crafting an item using a recipe.
//
// Added: v1.11.1
type CraftingEvent struct {
	// WindowID is the ID of the crafting window.
	//
	// Added: v1.11.1
	WindowID byte
	// CraftingType is the crafting source type. It is one of the CraftingEventType constants.
	//
	// Added: v1.11.1
	CraftingType int32
	// RecipeUUID is the UUID of the recipe that was crafted.
	//
	// Added: v1.11.1
	RecipeUUID uuid.UUID
	// Input is the list of input items used by the recipe.
	//
	// Added: v1.11.1
	Input []protocol.ItemInstance
	// Output is the list of output items created by the recipe.
	//
	// Added: v1.11.1
	Output []protocol.ItemInstance
}

// ID ...
func (*CraftingEvent) ID() uint32 {
	return IDCraftingEvent
}

func (pk *CraftingEvent) Marshal(io protocol.IO) {
	io.Uint8(&pk.WindowID)
	io.Varint32(&pk.CraftingType)
	io.UUID(&pk.RecipeUUID)
	protocol.FuncSlice(io, &pk.Input, io.ItemInstance)
	protocol.FuncSlice(io, &pk.Output, io.ItemInstance)
}
