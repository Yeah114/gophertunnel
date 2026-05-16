package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	GraphicsModeSimple = iota
	GraphicsModeFancy
	GraphicsModeAdvanced
	GraphicsModeRayTraced
)

// UpdateClientOptions is sent by the client when some of the client's options are updated, such as the
// graphics mode.
//
// Added: v1.21.70
type UpdateClientOptions struct {
	// GraphicsMode is the graphics mode that the client is using. It is one of the constants above.
	//
	// Added: v1.21.70
	GraphicsMode protocol.Optional[byte]
	// FilterProfanity is if the client only uses filtered messages or not.
	//
	// Added: v1.26.20
	FilterProfanity protocol.Optional[bool]
}

// ID ...
func (*UpdateClientOptions) ID() uint32 {
	return IDUpdateClientOptions
}

func (pk *UpdateClientOptions) Marshal(io protocol.IO) {
	protocol.OptionalFunc(io, &pk.GraphicsMode, io.Uint8)
	protocol.OptionalFunc(io, &pk.FilterProfanity, io.Bool)
}
