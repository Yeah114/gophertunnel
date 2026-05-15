package protocol

const (
	CreativeCategoryAll = iota
	CreativeCategoryConstruction
	CreativeCategoryNature
	CreativeCategoryEquipment
	CreativeCategoryItems
	CreativeCategoryItemCommandOnly
	CreativeCategoryUndefined
)

// CreativeGroup represents a group of items in the creative inventory. Each group has a category, name and an
// icon that represents the group.
//
// Added: v1.21.60
type CreativeGroup struct {
	// Category is the category the group falls under. It is one of the constants above.
	//
	// Added: v1.21.60
	Category int32
	// Name is the locale name of the group, i.e. "itemGroup.name.planks".
	//
	// Added: v1.21.60
	Name string
	// Icon is the item that represents the group in the creative inventory.
	//
	// Added: v1.21.60
	Icon ItemStack
}

// Marshal encodes/decodes a CreativeGroup.
func (x *CreativeGroup) Marshal(r IO) {
	r.Int32(&x.Category)
	r.String(&x.Name)
	r.Item(&x.Icon)
}

// CreativeItem represents a creative item present in the creative inventory.
//
// Added: v1.16
type CreativeItem struct {
	// CreativeItemNetworkID is a unique ID for the creative item. It has to be unique for each creative item
	// sent to the client. An incrementing ID per creative item does the job.
	//
	// Added: v1.16
	CreativeItemNetworkID uint32
	// Item is the item that should be added to the creative inventory.
	//
	// Added: v1.16
	Item ItemStack
	// GroupIndex is the index of the group that the item should be placed in. It is the index of the group in
	// the CreativeContent packet previously sent to the client.
	//
	// Added: v1.21.60
	GroupIndex uint32
}

// Marshal encodes/decodes a CreativeItem.
func (x *CreativeItem) Marshal(r IO) {
	r.Varuint32(&x.CreativeItemNetworkID)
	r.Item(&x.Item)
	r.Varuint32(&x.GroupIndex)
}
