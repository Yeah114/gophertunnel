package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// ClientBoundDataDrivenUIReload is sent by the server to reload the data-driven UI on the client.
//
// Added: v1.26.0
type ClientBoundDataDrivenUIReload struct{}

// ID ...
func (*ClientBoundDataDrivenUIReload) ID() uint32 {
	return IDClientBoundDataDrivenUIReload
}

func (pk *ClientBoundDataDrivenUIReload) Marshal(protocol.IO) {}
