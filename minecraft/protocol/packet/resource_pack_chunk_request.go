package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// ResourcePackChunkRequest is sent by the client to request a chunk of data from a particular resource pack,
// that it has obtained information about in a ResourcePackDataInfo packet.
//
// Added: v1.11.1
type ResourcePackChunkRequest struct {
	// UUID is the unique ID of the resource pack that the chunk of data is requested from.
	//
	// Added: v1.11.1
	UUID string
	// ChunkIndex is the requested chunk index of the chunk. It is a number that starts at 0 and is
	// incremented for each resource pack data chunk requested.
	//
	// Added: v1.11.1
	ChunkIndex uint32
}

// ID ...
func (*ResourcePackChunkRequest) ID() uint32 {
	return IDResourcePackChunkRequest
}

func (pk *ResourcePackChunkRequest) Marshal(io protocol.IO) {
	io.String(&pk.UUID)
	io.Uint32(&pk.ChunkIndex)
}
