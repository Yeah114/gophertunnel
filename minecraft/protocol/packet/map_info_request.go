package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// MapInfoRequest is sent by the client to request the server to deliver information of a certain map in the
// inventory of the player. The server should respond with a ClientBoundMapItemData packet.
//
// Added: v1.12
type MapInfoRequest struct {
	// MapID is the unique identifier that represents the map that is requested over network. It remains
	// consistent across sessions.
	//
	// Added: v1.12
	MapID int64
	// ClientPixels is a slice of pixels sent from the client to notify the server about the pixels that it isn't aware
	// of.
	//
	// Added: v1.19.20
	ClientPixels []protocol.PixelRequest
}

// ID ...
func (*MapInfoRequest) ID() uint32 {
	return IDMapInfoRequest
}

func (pk *MapInfoRequest) Marshal(io protocol.IO) {
	io.Varint64(&pk.MapID)
	if io.Protocol() >= protocol.Protocol1v19v20 {
		protocol.SliceUint32Length(io, &pk.ClientPixels)
	}
}
