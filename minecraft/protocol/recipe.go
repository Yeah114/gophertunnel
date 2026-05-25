package protocol

import (
	"github.com/google/uuid"
)

// PotionContainerChangeRecipe represents a recipe to turn a potion from one type to another. This means from
// a drinkable potion + gunpowder -> splash potion, and from a splash potion + dragon breath -> lingering
// potion.
//
// Added: v1.13.0
type PotionContainerChangeRecipe struct {
	// InputItemID is the item ID of the item to be put in. This is typically either the ID of a normal potion
	// or a splash potion.
	//
	// Added: v1.13.0
	InputItemID int32
	// ReagentItemID is the item ID of the item that needs to be added to the container in order to create the
	// output item.
	//
	// Added: v1.13.0
	ReagentItemID int32
	// OutputItemID is the item that is created using a combination of the InputItem and ReagentItem, which is
	// typically either the ID of a splash potion or a lingering potion.
	//
	// Added: v1.13.0
	OutputItemID int32
}

// Marshal encodes/decodes a PotionContainerChangeRecipe.
func (x *PotionContainerChangeRecipe) Marshal(r IO) {
	r.Varint32(&x.InputItemID)
	r.Varint32(&x.ReagentItemID)
	r.Varint32(&x.OutputItemID)
}

// PotionRecipe represents a potion mixing recipe which may be used in a brewing stand.
//
// Added: v1.13.0
type PotionRecipe struct {
	// InputPotionID is the item ID of the potion to be put in.
	//
	// Added: v1.13.0
	InputPotionID int32
	// InputPotionMetadata is the type of the potion to be put in. This is typically the meta of the
	// awkward potion (or water bottle to create an awkward potion).
	//
	// Added: v1.13.0
	InputPotionMetadata int32
	// ReagentItemID is the item ID of the item that needs to be added to the brewing stand in order to brew
	// the output potion.
	//
	// Added: v1.13.0
	ReagentItemID int32
	// ReagentItemMetadata is the metadata value of the item that needs to be added to the brewing stand in
	// order to brew the output potion.
	//
	// Added: v1.13.0
	ReagentItemMetadata int32
	// OutputPotionID is the item ID of the potion obtained as a result of the brewing recipe.
	//
	// Added: v1.13.0
	OutputPotionID int32
	// OutputPotionMetadata is the type of the potion that is obtained as a result of brewing the input
	// potion with the reagent item.
	//
	// Added: v1.13.0
	OutputPotionMetadata int32
}

// Marshal encodes/decodes a PotionRecipe.
func (x *PotionRecipe) Marshal(r IO) {
	r.Varint32(&x.InputPotionID)
	r.Varint32(&x.InputPotionMetadata)
	r.Varint32(&x.ReagentItemID)
	r.Varint32(&x.ReagentItemMetadata)
	r.Varint32(&x.OutputPotionID)
	r.Varint32(&x.OutputPotionMetadata)
}

const (
	RecipeUnlockContextNone = iota
	RecipeUnlockContextAlwaysUnlocked
	RecipeUnlockContextPlayerInWater
	RecipeUnlockContextPlayerHasManyItems
)

// RecipeUnlockRequirement represents a requirement that must be met in order to unlock a recipe. This is used
// for both shaped and shapeless recipes.
//
// Added: v1.21.0
type RecipeUnlockRequirement struct {
	// Context is the context in which the recipe is unlocked. This is one of the constants above.
	//
	// Added: v1.21.0
	Context byte
	// Ingredients are the ingredients required to unlock the recipe and only used if Context is set to none.
	//
	// Added: v1.21.0
	Ingredients []ItemDescriptorCount
}

// Marshal ...
func (x *RecipeUnlockRequirement) Marshal(r IO) {
	r.Uint8(&x.Context)
	if x.Context == RecipeUnlockContextNone {
		FuncSlice(r, &x.Ingredients, r.ItemDescriptorCount)
	}
}

const (
	RecipeShapeless int32 = iota
	RecipeShaped
	RecipeFurnace
	RecipeFurnaceData
	RecipeMulti
	RecipeShulkerBox
	RecipeShapelessChemistry
	RecipeShapedChemistry
	RecipeSmithingTransform
	RecipeSmithingTrim
)

// Recipe represents a recipe that may be sent in a CraftingData packet to let the client know what recipes
// are available server-side.
type Recipe interface {
	// Marshal encodes the recipe data to its binary representation into buf.
	Marshal(w *Writer)
	// Unmarshal decodes a serialised recipe from Reader r into the recipe instance.
	Unmarshal(r *Reader)
}

// lookupRecipe looks up the Recipe for a recipe type. False is returned if not
// found.
func lookupRecipe(recipeType int32, x *Recipe) bool {
	switch recipeType {
	case RecipeShapeless:
		*x = &ShapelessRecipe{}
	case RecipeShaped:
		*x = &ShapedRecipe{}
	case RecipeFurnace:
		*x = &FurnaceRecipe{}
	case RecipeFurnaceData:
		*x = &FurnaceDataRecipe{}
	case RecipeMulti:
		*x = &MultiRecipe{}
	case RecipeShulkerBox:
		*x = &ShulkerBoxRecipe{}
	case RecipeShapelessChemistry:
		*x = &ShapelessChemistryRecipe{}
	case RecipeShapedChemistry:
		*x = &ShapedChemistryRecipe{}
	case RecipeSmithingTransform:
		*x = &SmithingTransformRecipe{}
	case RecipeSmithingTrim:
		*x = &SmithingTrimRecipe{}
	default:
		return false
	}
	return true
}

// lookupRecipeType looks up the recipe type for a Recipe. False is returned if
// none was found.
func lookupRecipeType(x Recipe, recipeType *int32) bool {
	switch x.(type) {
	case *ShapelessRecipe:
		*recipeType = RecipeShapeless
	case *ShapedRecipe:
		*recipeType = RecipeShaped
	case *FurnaceRecipe:
		*recipeType = RecipeFurnace
	case *FurnaceDataRecipe:
		*recipeType = RecipeFurnaceData
	case *MultiRecipe:
		*recipeType = RecipeMulti
	case *ShulkerBoxRecipe:
		*recipeType = RecipeShulkerBox
	case *ShapelessChemistryRecipe:
		*recipeType = RecipeShapelessChemistry
	case *ShapedChemistryRecipe:
		*recipeType = RecipeShapedChemistry
	case *SmithingTransformRecipe:
		*recipeType = RecipeSmithingTransform
	case *SmithingTrimRecipe:
		*recipeType = RecipeSmithingTrim
	default:
		return false
	}
	return true
}

// ShapelessRecipe is a recipe that has no particular shape. Its functionality is shared with the
// RecipeShulkerBox and RecipeShapelessChemistry types.
//
// Added: v1.11.1
type ShapelessRecipe struct {
	// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes too,
	// but its functionality is not exactly known.
	//
	// Added: v1.19.60
	RecipeID string
	// Input is a list of items that serve as the input of the shapeless recipe. These items are the items
	// required to craft the output.
	//
	// Added: v1.11.1
	Input []ItemDescriptorCount
	// Output is a list of items that are created as a result of crafting the recipe.
	//
	// Added: v1.11.1
	Output []ItemStack
	// UUID is a UUID identifying the recipe. Since the CraftingEvent packet no longer exists, this can always be empty.
	//
	// Added: v1.11.1
	UUID uuid.UUID
	// Block is the block name that is required to craft the output of the recipe. The block is not prefixed
	// with 'minecraft:', so it will look like 'crafting_table' as an example.
	// The available blocks are:
	// - crafting_table
	// - cartography_table
	// - stonecutter
	// - furnace
	// - blast_furnace
	// - smoker
	// - campfire
	//
	// Added: v1.11.1
	Block string
	// Priority is the recipe priority used by the client when multiple recipes match the same input.
	//
	// Added: v1.11.1
	Priority int32
	// UnlockRequirement is a requirement that must be met in order to unlock the recipe.
	//
	// Added: v1.21.0
	UnlockRequirement RecipeUnlockRequirement
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	//
	// Added: v1.16.220
	RecipeNetworkID uint32
}

// ShulkerBoxRecipe is a shapeless recipe made specifically for shulker box crafting, so that they don't lose
// their user data when dyeing a shulker box.
//
// Added: v1.19.21
type ShulkerBoxRecipe struct {
	// ShapelessRecipe holds the underlying shapeless recipe data.
	//
	// Added: v1.19.21
	ShapelessRecipe
}

// ShapelessChemistryRecipe is a recipe specifically made for chemistry related features, which exist only in
// the Education Edition. They function the same as shapeless recipes do.
//
// Added: v1.19.21
type ShapelessChemistryRecipe struct {
	// ShapelessRecipe holds the underlying shapeless recipe data.
	//
	// Added: v1.19.21
	ShapelessRecipe
}

// ShapedRecipe is a recipe that has a specific shape that must be used to craft the output of the recipe.
// Trying to craft the item in any other shape will not work. The ShapedRecipe is of the same structure as the
// ShapedChemistryRecipe.
//
// Added: v1.11.1
type ShapedRecipe struct {
	// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes too,
	// but its functionality is not exactly known.
	//
	// Added: v1.19.60
	RecipeID string
	// Width is the width of the recipe's shape.
	//
	// Added: v1.11.1
	Width int32
	// Height is the height of the recipe's shape.
	//
	// Added: v1.11.1
	Height int32
	// Input is a list of items that serve as the input of the shapeless recipe. These items are the items
	// required to craft the output. The amount of input items must be exactly equal to Width * Height.
	//
	// Added: v1.11.1
	Input []ItemDescriptorCount
	// Output is a list of items that are created as a result of crafting the recipe.
	//
	// Added: v1.11.1
	Output []ItemStack
	// UUID is a UUID identifying the recipe. Since the CraftingEvent packet no longer exists, this can always be empty.
	//
	// Added: v1.11.1
	UUID uuid.UUID
	// Block is the block name that is required to craft the output of the recipe. The block is not prefixed
	// with 'minecraft:', so it will look like 'crafting_table' as an example.
	//
	// Added: v1.11.1
	Block string
	// Priority is the recipe priority used by the client when multiple recipes match the same input.
	//
	// Added: v1.11.1
	Priority int32
	// AssumeSymmetry specifies if the recipe is symmetrical. If this is set to true, the recipe will be
	// mirrored along the diagonal axis. This means that the recipe will be the same if rotated 180 degrees.
	//
	// Added: v1.20.80
	AssumeSymmetry bool
	// UnlockRequirement is a requirement that must be met in order to unlock the recipe.
	//
	// Added: v1.21.0
	UnlockRequirement RecipeUnlockRequirement
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	//
	// Added: v1.16.220
	RecipeNetworkID uint32
}

// ShapedChemistryRecipe is a recipe specifically made for chemistry related features, which exist only in the
// Education Edition. It functions the same as a normal ShapedRecipe.
//
// Added: v1.19.21
type ShapedChemistryRecipe struct {
	// ShapedRecipe holds the underlying shaped recipe data.
	//
	// Added: v1.19.21
	ShapedRecipe
}

// FurnaceRecipe is a recipe that is specifically used for all kinds of furnaces. These recipes don't just
// apply to furnaces, but also blast furnaces and smokers.
//
// Removed: v1.26.20
type FurnaceRecipe struct {
	// RecipeID is a unique ID of the recipe.
	//
	// Added: v1.26.20
	RecipeID string
	// InputType is the item type of the input item. The metadata value of the item is not used in the
	// FurnaceRecipe. Use FurnaceDataRecipe to allow an item with only one metadata value.
	InputType ItemType
	// Input is a list of items that serve as the input of the furnace recipe.
	//
	// Added: v1.26.20
	Input []ItemDescriptorCount
	// Output is the item that is created as a result of smelting/cooking an item in the furnace.
	Output ItemStack
	// UUID is a UUID identifying the recipe.
	//
	// Added: v1.26.20
	UUID uuid.UUID
	// Block is the block name that is required to create the output of the recipe. The block is not prefixed
	// with 'minecraft:', so it will look like 'furnace' as an example.
	Block string
	// Priority is the recipe priority used by the client when multiple recipes match the same input.
	//
	// Added: v1.26.20
	Priority int32
	// UnlockRequirement is a requirement that must be met in order to unlock the recipe.
	//
	// Added: v1.26.20
	UnlockRequirement RecipeUnlockRequirement
	// RecipeNetworkID is a unique ID used to identify the recipe over network.
	//
	// Added: v1.26.20
	RecipeNetworkID uint32
}

// FurnaceDataRecipe is a recipe specifically used for furnace-type crafting stations. It is equal to
// FurnaceRecipe, except it has an input item with a specific metadata value, instead of any metadata value.
//
// Removed: v1.26.20
type FurnaceDataRecipe struct {
	// FurnaceRecipe holds the underlying furnace recipe data.
	//
	// Removed: v1.26.20
	FurnaceRecipe
}

// MultiRecipe serves as an 'enable' switch for multi-shape recipes.
//
// Added: v1.11.1
type MultiRecipe struct {
	// UUID is a UUID identifying the recipe. Since the CraftingEvent packet no longer exists, this can always be empty.
	//
	// Added: v1.11.1
	UUID uuid.UUID
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	//
	// Added: v1.16.220
	RecipeNetworkID uint32
}

// SmithingTransformRecipe is a recipe specifically used for smithing tables. It has three input items and adds them
// together, resulting in a new item.
//
// Added: v1.19.60
type SmithingTransformRecipe struct {
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	//
	// Added: v1.19.60
	RecipeNetworkID uint32
	// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes too,
	// but its functionality is not exactly known.
	//
	// Added: v1.19.60
	RecipeID string
	// Template is the item that is used to shape the Base item based on the Addition being applied.
	//
	// Added: v1.19.60
	Template ItemDescriptorCount
	// Base is the item that the Addition is being applied to in the smithing table.
	//
	// Added: v1.19.60
	Base ItemDescriptorCount
	// Addition is the item that is being added to the Base item to result in a modified item.
	//
	// Added: v1.19.60
	Addition ItemDescriptorCount
	// Result is the resulting item from the two items being added together.
	//
	// Added: v1.19.60
	Result ItemStack
	// Block is the block name that is required to create the output of the recipe. The block is not prefixed with
	// 'minecraft:', so it will look like 'smithing_table' as an example.
	//
	// Added: v1.19.60
	Block string
}

// SmithingTrimRecipe is a recipe specifically used for applying armour trims to an armour piece inside a smithing table.
//
// Added: v1.19.80
type SmithingTrimRecipe struct {
	// RecipeNetworkID is a unique ID used to identify the recipe over network. Each recipe must have a unique
	// network ID. Recommended is to just increment a variable for each unique recipe registered.
	// This field must never be 0.
	//
	// Added: v1.19.80
	RecipeNetworkID uint32
	// RecipeID is a unique ID of the recipe. This ID must be unique amongst all other types of recipes too,
	// but its functionality is not exactly known.
	//
	// Added: v1.19.80
	RecipeID string
	// Template is the item that is used to shape the Base item based on the Addition being applied.
	//
	// Added: v1.19.80
	Template ItemDescriptorCount
	// Base is the item that the Addition is being applied to in the smithing table.
	//
	// Added: v1.19.80
	Base ItemDescriptorCount
	// Addition is the item that is being added to the Base item to result in a modified item.
	//
	// Added: v1.19.80
	Addition ItemDescriptorCount
	// Block is the block name that is required to create the output of the recipe. The block is not prefixed with
	// 'minecraft:', so it will look like 'smithing_table' as an example.
	//
	// Added: v1.19.80
	Block string
}

// Marshal ...
func (recipe *ShapelessRecipe) Marshal(w *Writer) {
	marshalShapeless(w, recipe)
}

// Unmarshal ...
func (recipe *ShapelessRecipe) Unmarshal(r *Reader) {
	marshalShapeless(r, recipe)
}

// Marshal ...
func (recipe *ShulkerBoxRecipe) Marshal(w *Writer) {
	marshalShapeless(w, &recipe.ShapelessRecipe)
}

// Unmarshal ...
func (recipe *ShulkerBoxRecipe) Unmarshal(r *Reader) {
	marshalShapeless(r, &recipe.ShapelessRecipe)
}

// Marshal ...
func (recipe *ShapelessChemistryRecipe) Marshal(w *Writer) {
	marshalShapeless(w, &recipe.ShapelessRecipe)
}

// Unmarshal ...
func (recipe *ShapelessChemistryRecipe) Unmarshal(r *Reader) {
	marshalShapeless(r, &recipe.ShapelessRecipe)
}

// Marshal ...
func (recipe *ShapedRecipe) Marshal(w *Writer) {
	marshalShaped(w, recipe)
}

// Unmarshal ...
func (recipe *ShapedRecipe) Unmarshal(r *Reader) {
	marshalShaped(r, recipe)
}

// Marshal ...
func (recipe *ShapedChemistryRecipe) Marshal(w *Writer) {
	marshalShaped(w, &recipe.ShapedRecipe)
}

// Unmarshal ...
func (recipe *ShapedChemistryRecipe) Unmarshal(r *Reader) {
	marshalShaped(r, &recipe.ShapedRecipe)
}

// Marshal ...
func (recipe *FurnaceRecipe) Marshal(w *Writer) {
	if w.Protocol() >= Protocol1v26v20 {
		marshalFurnaceShapeless(w, recipe)
		return
	}
	w.Varint32(&recipe.InputType.NetworkID)
	w.Item(&recipe.Output)
	w.String(&recipe.Block)
}

// Unmarshal ...
func (recipe *FurnaceRecipe) Unmarshal(r *Reader) {
	if r.Protocol() >= Protocol1v26v20 {
		marshalFurnaceShapeless(r, recipe)
		return
	}
	r.Varint32(&recipe.InputType.NetworkID)
	r.Item(&recipe.Output)
	r.String(&recipe.Block)
}

// Marshal ...
func (recipe *FurnaceDataRecipe) Marshal(w *Writer) {
	if w.Protocol() >= Protocol1v26v20 {
		marshalFurnaceShapeless(w, &recipe.FurnaceRecipe)
		return
	}
	w.Varint32(&recipe.InputType.NetworkID)
	aux := int32(recipe.InputType.MetadataValue)
	w.Varint32(&aux)
	w.Item(&recipe.Output)
	w.String(&recipe.Block)
}

// Unmarshal ...
func (recipe *FurnaceDataRecipe) Unmarshal(r *Reader) {
	if r.Protocol() >= Protocol1v26v20 {
		marshalFurnaceShapeless(r, &recipe.FurnaceRecipe)
		return
	}
	var dataValue int32
	r.Varint32(&recipe.InputType.NetworkID)
	r.Varint32(&dataValue)
	recipe.InputType.MetadataValue = uint32(dataValue)
	r.Item(&recipe.Output)
	r.String(&recipe.Block)
}

// Marshal ...
func (recipe *MultiRecipe) Marshal(w *Writer) {
	w.UUID(&recipe.UUID)
	w.Varuint32(&recipe.RecipeNetworkID)
}

// Unmarshal ...
func (recipe *MultiRecipe) Unmarshal(r *Reader) {
	r.UUID(&recipe.UUID)
	r.Varuint32(&recipe.RecipeNetworkID)
}

// Marshal ...
func (recipe *SmithingTransformRecipe) Marshal(w *Writer) {
	w.String(&recipe.RecipeID)
	w.ItemDescriptorCount(&recipe.Template)
	w.ItemDescriptorCount(&recipe.Base)
	w.ItemDescriptorCount(&recipe.Addition)
	w.Item(&recipe.Result)
	w.String(&recipe.Block)
	w.Varuint32(&recipe.RecipeNetworkID)
}

// Unmarshal ...
func (recipe *SmithingTransformRecipe) Unmarshal(r *Reader) {
	r.String(&recipe.RecipeID)
	r.ItemDescriptorCount(&recipe.Template)
	r.ItemDescriptorCount(&recipe.Base)
	r.ItemDescriptorCount(&recipe.Addition)
	r.Item(&recipe.Result)
	r.String(&recipe.Block)
	r.Varuint32(&recipe.RecipeNetworkID)
}

// Marshal ...
func (recipe *SmithingTrimRecipe) Marshal(w *Writer) {
	w.String(&recipe.RecipeID)
	w.ItemDescriptorCount(&recipe.Template)
	w.ItemDescriptorCount(&recipe.Base)
	w.ItemDescriptorCount(&recipe.Addition)
	w.String(&recipe.Block)
	w.Varuint32(&recipe.RecipeNetworkID)
}

// Unmarshal ...
func (recipe *SmithingTrimRecipe) Unmarshal(r *Reader) {
	r.String(&recipe.RecipeID)
	r.ItemDescriptorCount(&recipe.Template)
	r.ItemDescriptorCount(&recipe.Base)
	r.ItemDescriptorCount(&recipe.Addition)
	r.String(&recipe.Block)
	r.Varuint32(&recipe.RecipeNetworkID)
}

// marshalShaped ...
func marshalShaped(r IO, recipe *ShapedRecipe) {
	r.String(&recipe.RecipeID)
	r.Varint32(&recipe.Width)
	r.Varint32(&recipe.Height)
	FuncSliceOfLen(r, uint32(recipe.Width*recipe.Height), &recipe.Input, r.ItemDescriptorCount)
	FuncSlice(r, &recipe.Output, r.Item)
	r.UUID(&recipe.UUID)
	r.String(&recipe.Block)
	r.Varint32(&recipe.Priority)
	r.Bool(&recipe.AssumeSymmetry)
	Single(r, &recipe.UnlockRequirement)
	r.Varuint32(&recipe.RecipeNetworkID)
}

// marshalShapeless ...
func marshalShapeless(r IO, recipe *ShapelessRecipe) {
	r.String(&recipe.RecipeID)
	FuncSlice(r, &recipe.Input, r.ItemDescriptorCount)
	FuncSlice(r, &recipe.Output, r.Item)
	r.UUID(&recipe.UUID)
	r.String(&recipe.Block)
	r.Varint32(&recipe.Priority)
	Single(r, &recipe.UnlockRequirement)
	r.Varuint32(&recipe.RecipeNetworkID)
}

func marshalFurnaceShapeless(r IO, recipe *FurnaceRecipe) {
	r.String(&recipe.RecipeID)
	if _, ok := r.(*Writer); ok && len(recipe.Input) == 0 && recipe.InputType.NetworkID != 0 {
		recipe.Input = []ItemDescriptorCount{{
			Descriptor: &DefaultItemDescriptor{
				NetworkID:     int16(recipe.InputType.NetworkID),
				MetadataValue: int16(recipe.InputType.MetadataValue),
			},
			Count: 1,
		}}
	}
	FuncSlice(r, &recipe.Input, r.ItemDescriptorCount)
	if len(recipe.Input) > 0 {
		if input, ok := recipe.Input[0].Descriptor.(*DefaultItemDescriptor); ok {
			recipe.InputType = ItemType{NetworkID: int32(input.NetworkID), MetadataValue: uint32(input.MetadataValue)}
		}
	}
	output := []ItemStack{recipe.Output}
	FuncSlice(r, &output, r.Item)
	if len(output) > 0 {
		recipe.Output = output[0]
	}
	r.UUID(&recipe.UUID)
	r.String(&recipe.Block)
	r.Varint32(&recipe.Priority)
	Single(r, &recipe.UnlockRequirement)
	r.Varuint32(&recipe.RecipeNetworkID)
}
