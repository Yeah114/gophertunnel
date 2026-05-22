package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

// ServerStats is a packet sent from the server to the client to update the client on server statistics. It is purely
// used for telemetry.
//
// Added: v1.19.30
type ServerStats struct {
	// ServerTime ...
	//
	// Added: v1.19.30
	ServerTime float32
	// NetworkTime ...
	//
	// Added: v1.19.30
	NetworkTime float32
}

// ID ...
func (pk *ServerStats) ID() uint32 {
	return IDServerStats
}

func (pk *ServerStats) Marshal(io protocol.IO) {
	io.Float32(&pk.ServerTime)
	io.Float32(&pk.NetworkTime)
}
