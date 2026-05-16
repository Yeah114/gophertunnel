package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// EditorNetwork is a packet sent from the server to the client and vise-versa to communicate editor-mode related
// information. It carries a single compound tag containing the relevant information.
//
// Added: v1.19.10
type EditorNetwork struct {
	// RouteToManager ...
	//
	// Added: v1.21.20
	RouteToManager bool
	// Payload is a network little endian compound tag holding data relevant to the editor.
	//
	// Added: v1.19.10
	Payload map[string]any
}

// ID ...
func (*EditorNetwork) ID() uint32 {
	return IDEditorNetwork
}

func (pk *EditorNetwork) Marshal(io protocol.IO) {
	io.Bool(&pk.RouteToManager)
	io.NBT(&pk.Payload, nbt.NetworkLittleEndian)
}
