package protocol

import "github.com/sandertv/gophertunnel/minecraft/nbt"

const (
	ItemEntryVersionLegacy = iota
	ItemEntryVersionDataDriven
	ItemEntryVersionNone
)

// ItemInstance represents a unique instance of an item stack. These instances carry a specific network ID
// that is persistent for the stack.
//
// Added: v1.16
type ItemInstance struct {
	// StackNetworkID is the network ID of the item stack. If the stack is empty, 0 is always written for this
	// field. If not, the field should be set to 1 if the server authoritative inventories are disabled in the
	// StartGame packet, or to a unique stack ID if it is enabled.
	//
	// Added: v1.16
	StackNetworkID int32
	// Stack is the actual item stack of the item instance.
	//
	// Added: v1.16
	Stack ItemStack
}

// ItemStack represents an item instance/stack over network. It has a network ID and a metadata value that
// define its type.
//
// Added: v1.16
type ItemStack struct {
	ItemType
	// BlockRuntimeID is the runtime ID of the block represented by the item, if applicable.
	//
	// Added: v1.16
	BlockRuntimeID int32
	// Count is the count of items that the item stack holds.
	//
	// Added: v1.16
	Count uint16
	// NBTData is a map that is serialised to its NBT representation when sent in a packet.
	//
	// Added: v1.16
	NBTData map[string]any
	// CanBePlacedOn is a list of block identifiers like 'minecraft:stone' which the item, if it is an item
	// that can be placed, can be placed on top of.
	//
	// Added: v1.16
	CanBePlacedOn []string
	// CanBreak is a list of block identifiers like 'minecraft:dirt' that the item is able to break.
	//
	// Added: v1.16
	CanBreak []string
	// HasNetworkID specifies if StackNetworkID should be treated as present for this stack.
	//
	// Added: v1.16
	HasNetworkID bool
	// BlockingTick is the tick at which a shield started blocking. It is only used for shield items.
	//
	// Added: v1.16
	BlockingTick int64
}

// ItemType represents a consistent combination of network ID and metadata value of an item. It cannot usually
// be changed unless a new item is obtained.
//
// Added: v1.16
type ItemType struct {
	// NetworkID is the numerical network ID of the item. This is sometimes a positive ID, and sometimes a
	// negative ID, depending on what item it concerns.
	//
	// Added: v1.16
	NetworkID int32
	// MetadataValue is the metadata value of the item. For some items, this is the damage value, whereas for
	// other items it is simply an identifier of a variant of the item.
	//
	// Added: v1.16
	MetadataValue uint32
}

// ItemEntry is an item sent in the StartGame item table. It holds a name and a legacy ID, which is used to
// point back to that name.
//
// Added: v1.16
type ItemEntry struct {
	// Name if the name of the item, which is a name like 'minecraft:stick'.
	//
	// Added: v1.16
	Name string
	// RuntimeID is the ID that is used to identify the item over network. After sending all items in the
	// StartGame packet, items will then be identified using these numerical IDs.
	//
	// Added: v1.16
	RuntimeID int16
	// ComponentBased specifies if the item was created using components, meaning the item is a custom item.
	//
	// Added: v1.16
	ComponentBased bool
	// Version is the version of the item entry which is used by the client to determine how to handle the
	// item entry. It is one of the constants above.
	//
	// Added: v1.16
	Version int32
	// Data is a map containing the components and properties of the item, if the item is component based.
	//
	// Added: v1.16
	Data map[string]any
}

// Marshal encodes/decodes an ItemEntry.
func (x *ItemEntry) Marshal(r IO) {
	r.String(&x.Name)
	r.Int16(&x.RuntimeID)
	r.Bool(&x.ComponentBased)
	r.Varint32(&x.Version)
	r.NBT(&x.Data, nbt.NetworkLittleEndian)
}

// MaterialReducerOutput is an output from a material reducer.
//
// Added: v1.16
type MaterialReducerOutput struct {
	// NetworkID is the network ID of the output.
	//
	// Added: v1.16
	NetworkID int32
	// Count is the quantity of the output.
	//
	// Added: v1.16
	Count int32
}

// Marshal encodes/decodes a MaterialReducerOutput.
func (x *MaterialReducerOutput) Marshal(r IO) {
	r.Varint32(&x.NetworkID)
	r.Varint32(&x.Count)
}

// MaterialReducer is a craft in a material reducer block in education edition.
//
// Added: v1.16
type MaterialReducer struct {
	// InputItem is the starting item.
	//
	// Added: v1.16
	InputItem ItemType
	// Outputs contain all outputting items.
	//
	// Added: v1.16
	Outputs []MaterialReducerOutput
}
