package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// PlayerToggleCrafterSlotRequest is sent by the client when it tries to toggle the state of a slot within a Crafter.
//
// Added: v1.20.50
type PlayerToggleCrafterSlotRequest struct {
	// PosX is the X position of the Crafter that is being modified.
	//
	// Added: v1.20.50
	PosX int32
	// PosY is the Y position of the Crafter that is being modified.
	//
	// Added: v1.20.50
	PosY int32
	// PosZ is the Z position of the Crafter that is being modified.
	//
	// Added: v1.20.50
	PosZ int32
	// Slot is the index of the slot that was toggled. This should be a value between 0 and 8.
	//
	// Added: v1.20.50
	Slot byte
	// Disabled is the new state of the slot. If true, the slot is disabled, if false, the slot is enabled.
	//
	// Added: v1.20.50
	Disabled bool
}

// ID ...
func (*PlayerToggleCrafterSlotRequest) ID() uint32 {
	return IDPlayerToggleCrafterSlotRequest
}

func (pk *PlayerToggleCrafterSlotRequest) Marshal(io protocol.IO) {
	io.Int32(&pk.PosX)
	io.Int32(&pk.PosY)
	io.Int32(&pk.PosZ)
	io.Uint8(&pk.Slot)
	io.Bool(&pk.Disabled)
}
