package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	AnimateActionSwingArm = iota + 1
	_
	AnimateActionStopSleep
	AnimateActionCriticalHit
	AnimateActionMagicCriticalHit
)

const (
	AnimateActionRowRight = 128
	AnimateActionRowLeft  = 129
)

const (
	AnimateSwingSourceNone = iota
	AnimateSwingSourceBuild
	AnimateSwingSourceMine
	AnimateSwingSourceInteract
	AnimateSwingSourceAttack
	AnimateSwingSourceUseItem
	AnimateSwingSourceThrowItem
	AnimateSwingSourceDropItem
	AnimateSwingSourceEvent
)

// Animate is sent by the server to send a player animation from one player to all viewers of that player. It
// is used for a couple of actions, such as arm swimming and critical hits.
//
// Added: v1.12
type Animate struct {
	// ActionType is the ID of the animation action to execute. It is one of the action type constants that
	// may be found above.
	//
	// Added: v1.12
	// Changed: v1.21.130.28, encoded as a byte instead of a varint32.
	ActionType uint8
	// EntityRuntimeID is the runtime ID of the player that the animation should be played upon. The runtime
	// ID is unique for each world session, and entities are generally identified in packets using this
	// runtime ID.
	//
	// Added: v1.12
	EntityRuntimeID uint64
	// Data ...
	//
	// Added: v1.21.120
	Data float32
	// RowingTime is the rowing animation time for the row left and row right actions.
	//
	// Added: v1.12
	// Removed: v1.21.130.28
	RowingTime float32
	// SwingSource is the source for swing actions. It is one of the action type constants that
	// may be found above.
	//
	// Added: v1.21.130.28
	SwingSource uint8
}

// ID ...
func (*Animate) ID() uint32 {
	return IDAnimate
}

func (pk *Animate) Marshal(io protocol.IO) {
	if io.Protocol() >= protocol.Protocol1v21v130v28 {
		io.Uint8(&pk.ActionType)
	} else {
		actionType := int32(pk.ActionType)
		io.Varint32(&actionType)
		pk.ActionType = uint8(actionType)
	}
	io.Varuint64(&pk.EntityRuntimeID)
	if io.Protocol() >= protocol.Protocol1v21v120 {
		io.Float32(&pk.Data)
	}
	if io.Protocol() < protocol.Protocol1v21v130v28 && (pk.ActionType == AnimateActionRowRight || pk.ActionType == AnimateActionRowLeft) {
		io.Float32(&pk.RowingTime)
	}
	if io.Protocol() >= protocol.Protocol1v21v130v28 {
		var swingSource protocol.Optional[string]
		if pk.SwingSource != AnimateSwingSourceNone {
			swingSource = protocol.Option(swingSourceToString(pk.SwingSource))
		}
		protocol.OptionalFunc(io, &swingSource, io.String)
		if val, ok := swingSource.Value(); ok {
			swingSourceFromString(io, &pk.SwingSource, val)
		}
	}
}

func swingSourceFromString(io protocol.IO, x *uint8, s string) {
	switch s {
	case "none":
		*x = AnimateSwingSourceNone
	case "build":
		*x = AnimateSwingSourceBuild
	case "mine":
		*x = AnimateSwingSourceMine
	case "interact":
		*x = AnimateSwingSourceInteract
	case "attack":
		*x = AnimateSwingSourceAttack
	case "useitem":
		*x = AnimateSwingSourceUseItem
	case "throwitem":
		*x = AnimateSwingSourceThrowItem
	case "dropitem":
		*x = AnimateSwingSourceDropItem
	case "event":
		*x = AnimateSwingSourceEvent
	default:
		io.InvalidValue(s, "swingSource", "unknown source")
	}
}

func swingSourceToString(x uint8) string {
	switch x {
	case AnimateSwingSourceNone:
		return "none"
	case AnimateSwingSourceBuild:
		return "build"
	case AnimateSwingSourceMine:
		return "mine"
	case AnimateSwingSourceInteract:
		return "interact"
	case AnimateSwingSourceAttack:
		return "attack"
	case AnimateSwingSourceUseItem:
		return "useitem"
	case AnimateSwingSourceThrowItem:
		return "throwitem"
	case AnimateSwingSourceDropItem:
		return "dropitem"
	case AnimateSwingSourceEvent:
		return "event"
	default:
		return "unknown"
	}
}
