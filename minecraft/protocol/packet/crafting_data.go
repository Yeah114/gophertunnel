package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// CraftingData is sent by the server to let the client know all crafting data that the server maintains. This
// includes shapeless crafting, crafting table recipes, furnace recipes etc. Each crafting station's recipes
// are included in it.
//
// Added: v1.13
type CraftingData struct {
	// Recipes is a list of all recipes available on the server. It includes among others shapeless, shaped
	// and furnace recipes. The client will only be able to craft these recipes.
	//
	// Added: v1.13
	Recipes []protocol.Recipe
	// PotionRecipes is a list of all potion mixing recipes which may be used in the brewing stand.
	//
	// Added: v1.13.0
	PotionRecipes []protocol.PotionRecipe
	// PotionContainerChangeRecipes is a list of all recipes to convert a potion from one type to another,
	// such as from a drinkable potion to a splash potion, or from a splash potion to a lingering potion.
	//
	// Added: v1.13.0
	PotionContainerChangeRecipes []protocol.PotionContainerChangeRecipe
	// MaterialReducers is a list of all material reducers which is used in education edition chemistry.
	//
	// Added: v1.17.30
	MaterialReducers []protocol.MaterialReducer
	// ClearRecipes indicates if all recipes currently active on the client should be cleaned. Doing this
	// means that the client will have no recipes active by itself: Any CraftingData packets previously sent
	// will also be discarded, and only the recipes in this CraftingData packet will be used.
	//
	// Added: v1.13
	ClearRecipes bool
}

// ID ...
func (*CraftingData) ID() uint32 {
	return IDCraftingData
}

func (pk *CraftingData) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &pk.Recipes, io.Recipe)
	protocol.Slice(io, &pk.PotionRecipes)
	protocol.Slice(io, &pk.PotionContainerChangeRecipes)
	protocol.FuncSlice(io, &pk.MaterialReducers, io.MaterialReducer)
	io.Bool(&pk.ClearRecipes)
}
