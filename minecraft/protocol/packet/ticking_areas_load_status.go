package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// TickingAreasLoadStatus is sent by the server to the client to notify the client of a ticking area's loading status.
//
// Added: v1.18.30
type TickingAreasLoadStatus struct {
	// Preload is true if the server is waiting for the area's preload.
	//
	// Added: v1.18.30
	Preload bool
}

// ID ...
func (*TickingAreasLoadStatus) ID() uint32 {
	return IDTickingAreasLoadStatus
}

func (pk *TickingAreasLoadStatus) Marshal(io protocol.IO) {
	io.Bool(&pk.Preload)
}
