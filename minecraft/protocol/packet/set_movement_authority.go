package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

// SetMovementAuthority is sent by the server to change the movement authority mode.
type SetMovementAuthority struct {
	// MovementMode is one of the protocol.PlayerMovementMode* constants.
	MovementMode byte
}

// ID ...
func (*SetMovementAuthority) ID() uint32 {
	return IDSetMovementAuthority
}

func (pk *SetMovementAuthority) Marshal(io protocol.IO) {
	io.Uint8(&pk.MovementMode)
}
