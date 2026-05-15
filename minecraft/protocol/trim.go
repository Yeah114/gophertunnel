package protocol

// TrimPattern represents a pattern that can be applied to an armour piece in combination with a TrimMaterial.
// Added: v1.19.80
type TrimPattern struct {
	// ItemName is the identifier of the item that represents the pattern, for example
	// 'minecraft:wayfinder_armor_trim_smithing_template'.
	// Added: v1.19.80
	ItemName string
	// PatternID is the identifier of the pattern, for example, 'wayfinder'.
	// Added: v1.19.80
	PatternID string
}

// Marshal ...
func (x *TrimPattern) Marshal(r IO) {
	r.String(&x.ItemName)
	r.String(&x.PatternID)
}

// TrimMaterial represents a material that can be used when applying an armour trim.
// Added: v1.19.80
type TrimMaterial struct {
	// MaterialID is the identifier of the material, for example 'netherite'.
	// Added: v1.19.80
	MaterialID string
	// Colour is the colour code used for text formatting, for example '§j'.
	// Added: v1.19.80
	Colour string
	// ItemName is the identifier of the item that represents the material, for example, 'minecraft:netherite_ingot'.
	// Added: v1.19.80
	ItemName string
}

// Marshal ...
func (x *TrimMaterial) Marshal(r IO) {
	r.String(&x.MaterialID)
	r.String(&x.Colour)
	r.String(&x.ItemName)
}
