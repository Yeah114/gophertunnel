package protocol

// GameRule contains game rule data.
//
// Added: v1.17.0
type GameRule struct {
	// Name is the name of the game rule.
	//
	// Added: v1.17.0
	Name string
	// CanBeModifiedByPlayer specifies if the game rule can be modified by the player through the in-game UI.
	//
	// Added: v1.17.0
	CanBeModifiedByPlayer bool
	// Value is the new value of the game rule. This is either a bool, uint32 or float32.
	//
	// Added: v1.17.0
	Value any
}
