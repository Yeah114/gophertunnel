package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

// PlayerInput is sent by the client to send movement input in older input modes.
//
// Added: v1.11.1
type PlayerInput struct {
	// MotionX is the player's sideways movement input.
	//
	// Added: v1.11.1
	MotionX float32
	// MotionY is the player's forward movement input.
	//
	// Added: v1.11.1
	MotionY float32
	// Jumping specifies if the player is jumping.
	//
	// Added: v1.11.1
	Jumping bool
	// Sneaking specifies if the player is sneaking.
	//
	// Added: v1.11.1
	Sneaking bool
}

// ID ...
func (*PlayerInput) ID() uint32 {
	return IDPlayerInput
}

func (pk *PlayerInput) Marshal(io protocol.IO) {
	io.Float32(&pk.MotionX)
	io.Float32(&pk.MotionY)
	io.Bool(&pk.Jumping)
	io.Bool(&pk.Sneaking)
}
