package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	UnlockedRecipesTypeEmpty = iota
	UnlockedRecipesTypeInitiallyUnlocked
	UnlockedRecipesTypeNewlyUnlocked
	UnlockedRecipesTypeRemoveUnlocked
	UnlockedRecipesTypeRemoveAllUnlocked
)

// UnlockedRecipes gives the client a list of recipes that have been unlocked, restricting the recipes that appear in
// the recipe book.
//
// Added: v1.19.70
type UnlockedRecipes struct {
	// UnlockType is the type of unlock that the packet represents, and can either be adding or removing a list of recipes.
	// It is one of the constants listed above.
	//
	// Added: v1.19.70
	// Changed: v1.20.0.23, encoded as a little-endian uint32 action instead of a boolean newly unlocked flag.
	UnlockType uint32
	// Recipes is a list of recipe names that have been unlocked.
	//
	// Added: v1.19.70
	Recipes []string
}

// ID ...
func (*UnlockedRecipes) ID() uint32 {
	return IDUnlockedRecipes
}

func (pk *UnlockedRecipes) Marshal(io protocol.IO) {
	if io.Protocol() >= protocol.Protocol1v20v0v23 {
		io.Uint32(&pk.UnlockType)
	} else {
		newlyUnlocked := pk.UnlockType == UnlockedRecipesTypeNewlyUnlocked
		io.Bool(&newlyUnlocked)
		if newlyUnlocked {
			pk.UnlockType = UnlockedRecipesTypeNewlyUnlocked
		} else {
			pk.UnlockType = UnlockedRecipesTypeInitiallyUnlocked
		}
	}
	protocol.FuncSlice(io, &pk.Recipes, io.String)
}
