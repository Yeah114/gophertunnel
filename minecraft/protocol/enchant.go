package protocol

// EnchantmentOption represents a single option in the enchantment table for a single item.
//
// Added: v1.16
type EnchantmentOption struct {
	// Cost is the cost of the option. This is the amount of XP levels required to select this enchantment
	// option.
	//
	// Changed: v1.26.20.26, encoded as a uint8 instead of a varuint32.
	Cost uint8
	// Enchantments holds the enchantments that will be applied to the item when this option is clicked.
	//
	// Added: v1.16
	Enchantments ItemEnchantments
	// Name is a name that will be translated to the 'Standard Galactic Alphabet'
	// (https://minecraft.wiki/w/Enchanting_Table#Standard_Galactic_Alphabet) client-side. The names
	// generally have no meaning, such as:
	// 'animal imbue range galvanize '
	// 'bless inside creature shrink '
	// 'elder free of inside '
	//
	// Added: v1.16
	Name string
	// RecipeNetworkID is a unique network ID for this enchantment option. When enchanting, the client
	// will submit this network ID in a ItemStackRequest packet with the CraftRecipe action, so that the
	// server knows which enchantment was selected.
	// Note that this ID should still be unique with other actual recipes. It's recommended to start counting
	// for enchantment network IDs from the counter used for producing network IDs for the normal recipes.
	//
	// Added: v1.16
	RecipeNetworkID uint32
}

// Marshal encodes/decodes an EnchantmentOption.
func (x *EnchantmentOption) Marshal(r IO) {
	if r.Protocol() >= Protocol1v26v20v26 {
		r.Uint8(&x.Cost)
	} else {
		cost := uint32(x.Cost)
		r.Varuint32(&cost)
		x.Cost = uint8(cost)
	}
	Single(r, &x.Enchantments)
	r.String(&x.Name)
	r.Varuint32(&x.RecipeNetworkID)
}

// ItemEnchantments holds information on the enchantments that are applied to an item when a specific button
// is clicked in the enchantment table.
//
// Added: v1.16
type ItemEnchantments struct {
	// Slot is either 0, 1 or 2. Its exact usage is not clear.
	//
	// Added: v1.16
	Slot int32
	// Enchantments is an array of 3 slices of enchantment instances. Each array represents enchantments that
	// will be added to the item with a different activation type. The arrays in which enchantments are sent
	// by the vanilla server are as follows:
	// slice 1 { protection, fire protection, feather falling, blast protection, projectile protection,
	//           thorns, respiration, depth strider, aqua affinity, frost walker, soul speed }
	// slice 2 { sharpness, smite, bane of arthropods, fire aspect, looting, silk touch, unbreaking, fortune,
	//           flame, luck of the sea, impaling }
	// slice 3 { knockback, efficiency, power, punch, infinity, lure, mending, curse of binding,
	//           curse of vanishing, riptide, loyalty, channeling, multishot, piercing, quick charge }
	// The first slice holds armour enchantments, the differences between the slice 2 and slice 3 are more
	// vaguely defined.
	//
	// Added: v1.16
	Enchantments [3][]EnchantmentInstance
}

// Marshal encodes/decodes an ItemEnchantments.
func (x *ItemEnchantments) Marshal(r IO) {
	r.Int32(&x.Slot)
	for i := 0; i < 3; i++ {
		Slice(r, &x.Enchantments[i])
	}
}

// EnchantmentInstance represents a single enchantment instance with the type of the enchantment and its
// level.
//
// Added: v1.16
type EnchantmentInstance struct {
	// Type is the type of the enchantment.
	//
	// Changed: v1.26.20.26, encoded as a varuint32 instead of a uint8.
	Type byte
	// Level is the level of the enchantment.
	//
	// Added: v1.16
	Level byte
}

// Marshal encodes/decodes an EnchantmentInstance.
func (x *EnchantmentInstance) Marshal(r IO) {
	if r.Protocol() >= Protocol1v26v20v26 {
		typeID := uint32(x.Type)
		r.Varuint32(&typeID)
		x.Type = byte(typeID)
	} else {
		r.Uint8(&x.Type)
	}
	r.Uint8(&x.Level)
}
