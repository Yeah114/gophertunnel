package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

const (
	LoadingScreenTypeUnknown = iota
	LoadingScreenTypeStart
	LoadingScreenTypeEnd
)

// ServerBoundLoadingScreen is sent by the client to tell the server about the state of the loading
// screen that the client is currently displaying.
//
// Added: v1.21.20
type ServerBoundLoadingScreen struct {
	// Type is the type of the loading screen event. It is one of the constants that may be found above.
	//
	// Added: v1.21.20
	Type int32
	// LoadingScreenID is the ID of the screen that was previously sent by the server in the ChangeDimension
	// packet. The server should validate that the ID matches the last one it sent.
	//
	// Added: v1.21.20
	LoadingScreenID protocol.Optional[uint32]
}

// ID ...
func (*ServerBoundLoadingScreen) ID() uint32 {
	return IDServerBoundLoadingScreen
}

func (pk *ServerBoundLoadingScreen) Marshal(io protocol.IO) {
	io.Varint32(&pk.Type)
	protocol.OptionalFunc(io, &pk.LoadingScreenID, io.Uint32)
}
