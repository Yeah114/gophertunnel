package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// SetCommandsEnabled is sent by the server to enable or disable the ability to execute commands for the
// client. If disabled, the client itself will stop the execution of commands.
//
// Added: v1.11.1
type SetCommandsEnabled struct {
	// Enabled defines if the commands should be enabled, or if false, disabled.
	//
	// Added: v1.11.1
	Enabled bool
}

// ID ...
func (*SetCommandsEnabled) ID() uint32 {
	return IDSetCommandsEnabled
}

func (pk *SetCommandsEnabled) Marshal(io protocol.IO) {
	io.Bool(&pk.Enabled)
}
