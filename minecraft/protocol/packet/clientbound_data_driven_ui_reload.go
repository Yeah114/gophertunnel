package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ClientBoundDataDrivenUIReload is sent by the server to reload the data-driven UI on the client.
// This packet was added in v1.26.0.
type ClientBoundDataDrivenUIReload struct{}

// ID ...
func (*ClientBoundDataDrivenUIReload) ID() uint32 {
	return IDClientBoundDataDrivenUIReload
}

func (pk *ClientBoundDataDrivenUIReload) Marshal(protocol.IO) {}
