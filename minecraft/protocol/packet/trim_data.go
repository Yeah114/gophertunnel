package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// TrimData is sent by the server to the client when they first join the server.
//
// Added: v1.19.80
// It contains a list of all the patterns and materials that can be applied via armour trims.
type TrimData struct {
	// Patterns is a list of patterns that can be applied to armour. Each pattern has its own style and texture that is
	// defined through resource packs.
	//
	// Added: v1.19.80
	Patterns []protocol.TrimPattern
	// Materials is a list of materials that can be applied to armour. These are mostly different ores that have different
	// colours for more customization.
	//
	// Added: v1.19.80
	Materials []protocol.TrimMaterial
}

// ID ...
func (*TrimData) ID() uint32 {
	return IDTrimData
}

func (pk *TrimData) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Patterns)
	protocol.Slice(io, &pk.Materials)
}
