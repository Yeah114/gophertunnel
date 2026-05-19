package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// PhotoInfoRequest is sent by the client to request photo information from the server.
//
// Deprecated: This packet was deprecated in 1.19.80.
//
// Added: v1.17.40
type PhotoInfoRequest struct {
	// PhotoID is the ID of the photo.
	//
	// Added: v1.17.40
	PhotoID int64
}

// ID ...
func (*PhotoInfoRequest) ID() uint32 {
	return IDPhotoInfoRequest
}

func (pk *PhotoInfoRequest) Marshal(io protocol.IO) {
	io.Varint64(&pk.PhotoID)
}
