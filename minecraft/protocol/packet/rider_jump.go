package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

// RiderJump is sent by the client when jumping with a mount that supports charged jumps.
//
// Added: v1.11.1
type RiderJump struct {
	// JumpStrength is the charged jump strength, generally in the 0-100 range.
	//
	// Added: v1.11.1
	JumpStrength int32
}

// ID ...
func (*RiderJump) ID() uint32 {
	return IDRiderJump
}

func (pk *RiderJump) Marshal(io protocol.IO) {
	io.Varint32(&pk.JumpStrength)
}
