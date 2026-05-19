package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// ClientBoundDataDrivenUICloseScreen is sent by the server to close a data-driven UI screen on the client.
// This packet replaced ClientBoundDataDrivenUICloseAllScreens.
//
// Added: v1.26.10
// If FormID is not set, all data-driven UI screens are closed.
type ClientBoundDataDrivenUICloseScreen struct {
	// FormID is the optional unique instance ID of the form to close. If not set, all forms are closed.
	//
	// Added: v1.26.10
	FormID protocol.Optional[uint32]
}

// ID ...
func (*ClientBoundDataDrivenUICloseScreen) ID() uint32 {
	return IDClientBoundDataDrivenUICloseScreen
}

func (pk *ClientBoundDataDrivenUICloseScreen) Marshal(io protocol.IO) {
	if io.Protocol() >= protocol.Protocol1v26v10 {
		protocol.OptionalFunc(io, &pk.FormID, io.Uint32)
	}
}
