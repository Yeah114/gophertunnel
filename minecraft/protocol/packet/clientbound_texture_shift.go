package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	TextureShiftActionInvalid = iota
	TextureShiftActionInitialize
	TextureShiftActionStart
	TextureShiftActionSetEnabled
	TextureShiftActionSync
)

// ClientBoundTextureShift is sent by the server to control texture shift animations on the client.
//
// Added: v1.26.0
type ClientBoundTextureShift struct {
	// ActionID is the texture shift action to perform. It is one of the constants above.
	//
	// Added: v1.26.0
	ActionID uint8
	// CollectionName is the name of the texture shift collection.
	//
	// Added: v1.26.0
	CollectionName string
	// FromStep is the step to shift from.
	//
	// Added: v1.26.0
	FromStep string
	// ToStep is the step to shift to.
	//
	// Added: v1.26.0
	ToStep string
	// AllSteps is a list of all steps in the texture shift.
	//
	// Added: v1.26.0
	AllSteps []string
	// CurrentLengthTicks is the current length of the shift in ticks.
	//
	// Added: v1.26.0
	CurrentLengthTicks uint64
	// TotalLengthTicks is the total length of the shift in ticks.
	//
	// Added: v1.26.0
	TotalLengthTicks uint64
	// Enabled specifies if the texture shift is enabled.
	//
	// Added: v1.26.0
	Enabled bool
}

// ID ...
func (*ClientBoundTextureShift) ID() uint32 {
	return IDClientBoundTextureShift
}

func (pk *ClientBoundTextureShift) Marshal(io protocol.IO) {
	io.Uint8(&pk.ActionID)
	io.String(&pk.CollectionName)
	io.String(&pk.FromStep)
	io.String(&pk.ToStep)
	protocol.FuncSlice(io, &pk.AllSteps, io.String)
	io.Varuint64(&pk.CurrentLengthTicks)
	io.Varuint64(&pk.TotalLengthTicks)
	io.Bool(&pk.Enabled)
}
